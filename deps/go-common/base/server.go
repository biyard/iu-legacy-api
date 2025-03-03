package base

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"biyard.co/common/logger"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type Task func() *Error

var (
	startWebWorker = func(r *Server) error {
		// w := &responseWrite{}

		// body := `{"password":"","hint":""}`

		// req, e := http.NewRequest("POST", "/v1/wallet", bytes.NewBuffer([]byte(body)))
		// if e != nil {
		// 	return e
		// }

		// r.g.ServeHTTP(w, req)

		return nil
	}
	WatchFiles = func(*Server) {}
)

// Server provides a server wrapper to serve http
type Server struct {
	g      *gin.Engine
	cfg    *Config
	ctx    context.Context
	cancel context.CancelFunc
	cron   *cron.Cron

	srv     *http.Server
	watcher *fsnotify.Watcher
	adapter *ginadapter.GinLambda
	tasks   []Task
}

func NewServer(conf Config) *Server {
	var g *gin.Engine
	var c *cron.Cron

	if conf.Server.Type != ServerTypeLambdaTask {
		gin.SetMode(gin.ReleaseMode)
		g = gin.New()
		c = cron.New()

		g.Use(gin.Recovery())
		g.Use(corsHandler(conf))
		g.Use(logHandler())
		logger.Init(conf.Logging)
	}
	ctx, cancel := context.WithCancel(context.Background())

	ret := &Server{
		g:      g,
		cfg:    &conf,
		ctx:    ctx,
		cron:   c,
		cancel: cancel,
	}

	return ret
}

func (r *Server) Context() context.Context {
	return r.ctx
}

func (r *Server) RunTask(ctx context.Context, event interface{}) error {
	for _, task := range r.tasks {
		err := task()
		if err != nil {
			return err
		}
	}
	return nil
}

// LambdaIntegratedProxy is a handler for lambda integrated proxy.
// It can handle both APIGateway and EventBridge events.
// This function is under incubating so be careful to use it in production stage.
func (r *Server) LambdaIntegratedProxy(ctx context.Context, req IntegratedEvent) (interface{}, error) {
	switch req.Type() {
	case APIGateway:
		return r.adapter.ProxyWithContext(ctx, req.APIGatewayProxyRequest)
	case EventBridge:
		err := r.RunTask(ctx, req.EventBridgeEvent)
		return nil, err
	default:
		logger.Errorf(ctx, "Unknown event type: %s", req.Type)
	}

	return nil, ErrInvalidLambdaType
}

func (r *Server) Start() error {
	if r.cfg.Server.Type == ServerTypeWebWorker {
		startWebWorker(r)
		return nil
	}
	if r.cfg.Server.Type == ServerTypeLambda {
		logger.Infof(nil, "Starting server on lambda")
		r.adapter = ginadapter.New(r.g)

		lambda.Start(r.adapter.ProxyWithContext)
		return nil
	} else if r.cfg.Server.Type == ServerTypeLambdaTask {
		// NOTE: This is a experimental code for supporting both APIGateway and EventBridge events.
		lambda.Start(r.LambdaIntegratedProxy)

		return nil
	}

	logger.Infof(nil, "Starting server on port %d", r.cfg.Server.Port)
	port := r.cfg.Server.Port
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r.g,
	}
	r.srv = srv
	if r.cfg.Watch.Enabled {
		WatchFiles(r)
	}
	r.cron.Start()

	err := srv.ListenAndServe()
	for err != nil && strings.Contains(err.Error(), "bind: address already in use") {
		port = port + 1
		srv.Addr = fmt.Sprintf(":%d", port)
		logger.Infof(nil, "Re-trying server on port %d", port)
		err = srv.ListenAndServe()
	}

	return err
}

func (r *Server) Stop() {
	r.srv.Shutdown(r.ctx)
	r.cancel()
	r.watcher.Close()
}

// Filter filters directories we don't need to watch such like .git
func (r *Server) Filter(path string) bool {
	filter := []string{".git", "node_modules", "vendor", "logs", "bin", ".build"}
	for _, f := range filter {
		if path == f {
			return false
		}
	}

	return true
}

func (r *Server) Use(middlewares ...gin.HandlerFunc) {
	r.g.Use(middlewares...)
}

func (r *Server) Group(relativePath string, routers ...Router) {
	rg := r.g.Group(relativePath)

	for _, router := range routers {
		router.OnInit(NewRoute(rg))
	}
}

func (r *Server) Schedule(schedule string, handler Scheduler) {
	if r.cron != nil {
		r.cron.AddFunc(schedule, func() {
			err := handler.RunTask()
			if err != nil {
				logger.Errorf(nil, "Error on schedule: %s", err.Error())
			}
		})
		handler.RunTask()
	} else {
		r.tasks = append(r.tasks, handler.RunTask)
	}
}

func corsHandler(conf Config) gin.HandlerFunc {
	if conf.CORS.Enabled {
		return cors.New(cors.Config{
			AllowOrigins:     conf.CORS.Origins,
			AllowHeaders:     conf.CORS.Headers,
			AllowAllOrigins:  true,
			AllowMethods:     conf.CORS.Methods,
			AllowCredentials: conf.CORS.CredentialMode,
		})
	}

	return cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{
			// General headers
			"Origin", "Content-Length", "Content-Type", "Authorization",
			"Accept", "Accept-Encoding", "Accept-Language", "Connection",
			"Host", "Referer", "User-Agent", "Cache-Control", "Pragma", "Expires",

			// Client tracing
			"X-Real-IP", "X-Forwarded-For", "X-Forwarded-Proto", "X-Forwarded-Host", "X-Forwarded-Port",
			"X-Auth-Token", "X-CSRF-Token", "X-Requested-With", "If-Modified-Since", "If-None-Match", "Last-Modified", "Content-Disposition", "Content-Transfer-Encoding",

			// Opentracing headers
			"X-Request-Id", "X-B3-TraceId", "X-B3-SpanId", "X-B3-ParentSpanId", "X-B3-Sampled", "X-B3-Flags",

			// AWS API Gateway headers
			"X-Amz-Date", "X-Api-Key", "X-Amz-Security-Token",
		},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	})
}

func logHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		// Process request
		c.Next()

		// Log only when path is not being skipped
		// if _, ok := skip[path]; !ok {
		// Stop timer
		timestamp := time.Now()
		latency := timestamp.Sub(start)
		bodySize := c.Writer.Size()

		msg := map[string]any{
			"timestamp":  timestamp,
			"method":     c.Request.Method,
			"path":       path,
			"statusCode": c.Writer.Status(),
			"clientIp":   c.ClientIP(),
			"latency":    latency,
			"bodySize":   bodySize,
		}

		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		if errorMessage != "" {
			msg["error"] = errorMessage
		}

		logger.Infos(NewBaseContextFromGin(c), msg)
	}
}
