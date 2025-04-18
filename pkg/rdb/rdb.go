package rdb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sysatom/framework/pkg/config"
	"go.uber.org/fx"
)

func NewRedisClient(lc fx.Lifecycle) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", config.App.Redis.Host, config.App.Redis.Port)
	password := config.App.Redis.Password
	if addr == ":" || password == "" {
		return nil, errors.New("redis config error")
	}
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           config.App.Redis.DB,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			s := client.Ping(context.Background())
			_, err := s.Result()
			if err != nil {
				return fmt.Errorf("redis server error %w", err)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return client, nil
}
