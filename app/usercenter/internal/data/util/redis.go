package util

import (
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"kratos-admin/app/usercenter/internal/conf"
)

func NewRedis(c *conf.Data_Redis) (*redis.Client, error) {
	options := &redis.Options{
		Addr: c.Addr,
		DB:   int(c.Db),
	}
	if c.Auth != "" {
		options.Password = c.Auth
	}
	if c.MaxConn > 0 {
		options.PoolSize = int(c.MaxConn)
	}
	if c.MaxIdle > 0 {
		options.MinIdleConns = int(c.MaxIdle)
	}
	if c.MaxRetry > 0 {
		options.MaxRetries = int(c.MaxRetry)
	}
	if c.ReadTimeout != nil {
		options.ReadTimeout = c.ReadTimeout.AsDuration()
	}
	if c.WriteTimeout != nil {
		options.WriteTimeout = c.WriteTimeout.AsDuration()
	}
	rds := redis.NewClient(options)
	if err := redisotel.InstrumentTracing(rds); err != nil {
		_ = rds.Close()
		return nil, err
	}
	return rds, nil
}
