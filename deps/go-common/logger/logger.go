package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type M map[string]interface{}

var (
	DebugH = func(ctx context.Context, h func()) {}
	Debug  = func(ctx context.Context, args ...interface{}) {}
	Debugs = func(ctx context.Context, args interface{}) {}
	Debugf = func(ctx context.Context, format string, args ...interface{}) {}

	Info  = Debug
	Infof = Debugf
	Infos = Debugs

	Warn  = Debug
	Warnf = Debugf
	Warns = Debugs

	Error  = Debug
	Errorf = Debugf
	Errors = Debugs

	Critical  = Debug
	Criticalf = Debugf
	Criticals = Debugs

	Audit  = manLog(zapcore.Level(AuditLevel))
	Auditf = manLogf(zapcore.Level(AuditLevel))
	Audits = manLogs(zapcore.Level(AuditLevel))

	ManipulateLogs = func(ctx LoggableContext, res interface{}) interface{} { return res }
	addLog         = func(ctx LoggableContext, msg []interface{}) {}
	mu             sync.RWMutex
)

var reqIDs = map[uint64]string{}
var rMu sync.RWMutex
var GlobalLogger *zap.Logger

const loggerCallStackDepth = 5

type DebugLog struct {
	Stack   string        `json:"stack"`
	Message []interface{} `json:"message"`
}

type LoggingConfig struct {
	Level    string `json:"level" yaml:"level"`
	Encoding string `json:"encoding" yaml:"encoding"`

	LogRotate LogRotateConfig `yaml:"logRotate"`
}

type LogRotateConfig struct {
	Enabled      bool   `json:"enabled" yaml:"enabled"`
	LogPath      string `json:"logPath" yaml:"logPath"`
	ErrorLogPath string `json:"errorLogPath" yaml:"errorLogPath"`
	MaxSize      int    `json:"maxSize" yaml:"maxSize"`
	MaxBackups   int    `json:"maxBackups" yaml:"maxBackups"`
	MaxAge       int    `json:"maxAge" yaml:"maxAge"`
}

func Init(c LoggingConfig) *zap.Logger {
	level := c.Level

	var syncers []zapcore.WriteSyncer
	var errSyncer zapcore.WriteSyncer

	s, closer, err := zap.Open("stdout")
	if err != nil {
		closer()
	} else {
		syncers = append(syncers, s)
	}

	if c.LogRotate.Enabled {
		if c.LogRotate.LogPath != "" {
			syncers = append(syncers, zapcore.AddSync(&lumberjack.Logger{
				Filename:   c.LogRotate.LogPath,
				MaxSize:    c.LogRotate.MaxSize, // megabytes
				MaxBackups: c.LogRotate.MaxBackups,
				MaxAge:     c.LogRotate.MaxAge, // days
			}))
		}

		if c.LogRotate.ErrorLogPath == "" {
			c.LogRotate.ErrorLogPath = "stderr"
		}

		errSyncer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   c.LogRotate.ErrorLogPath,
			MaxSize:    c.LogRotate.MaxSize, // megabytes
			MaxBackups: c.LogRotate.MaxBackups,
			MaxAge:     c.LogRotate.MaxAge, // days
		})
	}
	logWriter := zap.CombineWriteSyncers(syncers...)
	reqIDs = map[uint64]string{}

	if level == "debug" {
		ManipulateLogs = manipulateLogs
		DebugH = func(ctx context.Context, h func()) { h() }
		addLog = _addLog
	}

	_encoderConstructor := map[string]func(zapcore.EncoderConfig) zapcore.Encoder{
		"console": zapcore.NewConsoleEncoder,
		"json":    NewJSONEncoder,
	}

	enc := _encoderConstructor[c.Encoding](zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:       "T",
		LevelKey:      "L",
		NameKey:       "N",
		CallerKey:     "C",
		MessageKey:    "M",
		StacktraceKey: "S",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   LevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02T15:04:05.000Z1600"))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			c := zapcore.NewEntryCaller(runtime.Caller(loggerCallStackDepth))
			log := c.FullPath()
			l := strings.Split(log, "kms-klip/")
			ll := len(l)

			if ll > 1 {
				log = l[ll-1]
			}
			enc.AppendString(log)
		},
	})

	logEnablers := map[string]zapcore.LevelEnabler{
		"critical": zapcore.PanicLevel,
		"error":    zapcore.ErrorLevel,
		"warning":  zapcore.WarnLevel,
		"info":     zapcore.InfoLevel,
		"debug":    zapcore.DebugLevel,
	}
	hostname, _ := os.Hostname()

	logger := zap.New(zapcore.NewCore(enc, logWriter, logEnablers[level]),
		zap.ErrorOutput(errSyncer),
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel)).Named(hostname)
	zap.ReplaceGlobals(logger)
	logLevel := map[string]int{"critical": 1, "error": 2, "warning": 3, "info": 4, "debug": 5}
	logHandler := []*func(context.Context, ...interface{}){&Critical, &Error, &Warn, &Info, &Debug}
	logHandlerf := []*func(context.Context, string, ...interface{}){&Criticalf, &Errorf, &Warnf, &Infof, &Debugf}
	logHandlers := []*func(context.Context, interface{}){&Criticals, &Errors, &Warns, &Infos, &Debugs}
	lvls := []zapcore.Level{zapcore.PanicLevel, zapcore.ErrorLevel, zapcore.WarnLevel, zapcore.InfoLevel, zapcore.DebugLevel}

	for i := 0; i < logLevel[level]; i++ {
		*logHandler[i] = manLog(lvls[i])
		*logHandlers[i] = manLogs(lvls[i])
		*logHandlerf[i] = manLogf(lvls[i])
	}

	for i := logLevel[level]; i < len(logHandler); i++ {
		*logHandler[i] = func(ctx context.Context, args ...interface{}) {}

		*logHandlerf[i] = func(ctx context.Context, format string, args ...interface{}) {}
	}

	GlobalLogger = logger

	return logger
}

type objecter interface {
	Object() any
}

func manipulateArgs(args []interface{}) []interface{} {
	ret := make([]interface{}, len(args))
	for i := 0; i < len(args); i++ {
		if v, ok := args[i].(objecter); ok {
			ret[i] = v.Object()
		} else if v, ok := args[i].(error); ok {
			ret[i] = v.Error()
		} else {
			ret[i] = args[i]
		}
	}

	return ret
}

func manipulateLogs(ctx LoggableContext, res interface{}) interface{} {
	var m interface{} = map[string]interface{}{}

	if kind := reflect.Indirect(reflect.ValueOf(res)).Kind(); kind != reflect.Struct && kind != reflect.Map {
		m = []interface{}{}
	}

	doc, _ := json.Marshal(res)
	json.Unmarshal(doc, &m)
	switch v := m.(type) {
	case []interface{}:
		m = append(v, ctx.Logs())
	case map[string]interface{}:
		v["zlogs"] = ctx.Logs()
	}

	return m
}

func _addLog(ctx LoggableContext, msg []interface{}) {
	stack := strings.Split(string(debug.Stack()), "\n")[8]
	stack = strings.ReplaceAll(stack, "\t", "")
	for i := 0; i < len(msg); i++ {
		if reflect.Indirect(reflect.ValueOf(msg[i])).Kind() == reflect.Func {
			msg[i] = fmt.Sprintf("%v", msg[i])
		}
	}
	ctx.AddLog(stack, msg)
}

func manLog(lvl zapcore.Level) func(ctx context.Context, args ...interface{}) {
	return func(ctx context.Context, args ...interface{}) {
		DebugH(ctx, func() {
			args = append([]interface{}{time.Now().Format("2006-01-02 15:04:05.000000")}, args...)
		})

		args = manipulateArgs(args)
		if c, ok := ctx.(LoggableContext); ok {
			addLog(c, args)
		}
		form := ""
		msg, values := newLog(ctx, form, args...)
		if ce := GlobalLogger.Check(lvl, msg); ce != nil {
			ce.Write(values...)
		}
	}
}

func manLogf(lvl zapcore.Level) func(ctx context.Context, format string, args ...interface{}) {
	return func(ctx context.Context, format string, args ...interface{}) {
		DebugH(ctx, func() {
			format = "[%v] " + format
			args = append([]interface{}{time.Now().Format("2006-01-02 15:04:05.000000")}, args...)
		})

		args = manipulateArgs(args)
		l := append([]interface{}{fmt.Sprintf(format, args...)}, args...)
		if c, ok := ctx.(LoggableContext); ok {
			addLog(c, l)
		}
		msg, values := newLog(ctx, format, args...)
		if ce := GlobalLogger.Check(lvl, msg); ce != nil {
			ce.Write(values...)
		}
	}
}

func manLogs(lvl zapcore.Level) func(ctx context.Context, arg interface{}) {
	return func(ctx context.Context, arg interface{}) {
		if c, ok := ctx.(LoggableContext); ok {
			addLog(c, []interface{}{arg})
		}
		msg, values := newLogForStruct(ctx, arg)
		if ce := GlobalLogger.Check(lvl, msg); ce != nil {
			ce.Write(values...)
		}
	}
}
