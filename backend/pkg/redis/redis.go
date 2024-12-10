package redis

import (
	"EduConnect/pkg/logger"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	URI string `mapstructure:"uri"`
}

type RedisConnector struct {
	log logger.Logger
	cfg *RedisConfig
}

func NewRedisConnector(log logger.Logger, cfg *RedisConfig) *RedisConnector {
	return &RedisConnector{log: log, cfg: cfg}
}

func (r *RedisConnector) NewRedisConn(ctx context.Context) (*redis.Client, error) {
	rsc := redis.NewClient(&redis.Options{
		Addr: r.cfg.URI,
	})

	if err := rsc.Ping(ctx).Err(); err != nil {
		r.log.Error("(RedisConnector) error: ", err)
		return nil, fmt.Errorf("failed to ping redis: %v", err)
	}

	return rsc, nil
}
