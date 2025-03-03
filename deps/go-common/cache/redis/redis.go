package redis

import (
	"context"
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"github.com/go-redsync/redsync/v4"
	goredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

const lockPrefix = "locker"

var once sync.Once
var db *Redis

type Redis struct {
	client  redis.UniversalClient
	sync    *redsync.Redsync
	timeout time.Duration
	ctx     context.Context
	cfg     base.RedisConfig
}

func NewRedis(ctx context.Context, conf base.RedisConfig) (*Redis, error) {
	e := (*base.Error)(nil)

	once.Do(func() {
		ts, err := time.ParseDuration(conf.PoolTimeout)
		if err != nil {
			logger.Errorf(ctx, "error parsing pool timeout: %v %v", conf.PoolTimeout, err)
			e = base.ErrRedisConfigParse
			return
		}
		rts, err := time.ParseDuration(conf.ReadTimeout)
		if err != nil {
			logger.Errorf(ctx, "error parsing read timeout: %v %v", conf.ReadTimeout, err)
			e = base.ErrRedisConfigParse
			return
		}
		wts, err := time.ParseDuration(conf.WriteTimeout)
		if err != nil {
			logger.Errorf(ctx, "error parsing write timeout: %v %v", conf.WriteTimeout, err)
			e = base.ErrRedisConfigParse
			return
		}

		minrts, err := time.ParseDuration(conf.MinRetryBackoff)
		if err != nil {
			logger.Errorf(ctx, "error parsing min retry backoff: %v %v", conf.MinRetryBackoff, err)
			e = base.ErrRedisConfigParse
			return
		}

		maxrts, err := time.ParseDuration(conf.MaxRetryBackoff)
		if err != nil {
			logger.Errorf(ctx, "error parsing max retry backoff: %v %v", conf.MaxRetryBackoff, err)
			e = base.ErrRedisConfigParse
			return
		}

		c := &redis.ClusterOptions{
			Addrs:           []string{conf.Endpoint},
			PoolSize:        conf.PoolSize,
			PoolTimeout:     ts,
			ReadTimeout:     rts,
			WriteTimeout:    wts,
			MaxRetries:      conf.MaxRetries,
			MinRetryBackoff: minrts,
			MaxRetryBackoff: maxrts,
			ReadOnly:        conf.ReadOnly,
			RouteRandomly:   conf.RandomRoute,
			MinIdleConns:    conf.MinIdleConn,
		}
		if conf.TLS {
			c.TLSConfig = &tls.Config{
				MinVersion: tls.VersionTLS12,
			}
		}

		var cli redis.UniversalClient

		cli = redis.NewClusterClient(c)
		if conf.Single {
			cli = redis.NewClient(&redis.Options{
				Addr:         conf.Endpoint,
				PoolSize:     conf.PoolSize,
				PoolTimeout:  ts,
				ReadTimeout:  rts,
				WriteTimeout: wts,
				MaxRetries:   conf.MaxRetries,
				MinIdleConns: conf.MinIdleConn,
			})
		}

		timeout, err := time.ParseDuration(conf.CacheExpiration)
		if err != nil {
			logger.Errorf(ctx, "error parsing redis cache expiration: %v %v", conf.CacheExpiration, err)
			e = base.ErrRedisConfigParse
			return
		}
		pool := goredis.NewPool(cli)

		if err := cli.Ping(ctx).Err(); err != nil {
			logger.Errorf(ctx, "redis ping error: %v", err)
		}

		db = &Redis{
			client:  cli,
			sync:    redsync.New(pool),
			timeout: timeout,
			ctx:     ctx,
			cfg:     conf,
		}
	})

	return db, e
}

func (r *Redis) GetClient() redis.UniversalClient {
	return r.client
}

func (r *Redis) GetSync() *redsync.Redsync {
	return r.sync
}

func (r *Redis) NewLock(key string) sync.Locker {
	return NewLock(r.ctx,
		r.sync.NewMutex(r.Realkey(fmt.Sprintf("%v##%v", lockPrefix, key)), redsync.WithExpiry(r.timeout)))
}

func (r *Redis) Realkey(key string) string {
	return fmt.Sprintf("%v##%v", r.cfg.ServiceID, key)
}

func (r *Redis) Cache(ctx context.Context, key string, value string) error {
	if err := r.client.Set(ctx, r.Realkey(key), value, r.timeout).Err(); err != nil {
		logger.Errorf(ctx, "error caching document: %v", err)
		return base.ErrCachePush
	}

	return nil
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, r.Realkey(key)).Result()
	if err != nil {
		return "", base.ErrCacheGet
	}

	return val, nil
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	if err := r.client.Del(ctx, r.Realkey(key)).Err(); err != nil {
		logger.Errorf(ctx, "error deleting document: %v", err)
		return base.ErrCacheDelete
	}

	return nil
}
