package repository

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"github.com/sysatom/framework/ent"
	"github.com/sysatom/framework/pkg/config"
	"go.uber.org/fx"
	"time"
)

func NewMySQLClient(lc fx.Lifecycle, _ config.Type) (*ent.Client, error) {
	client, err := ent.Open("mysql", config.App.Store.MySQL.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// Run the auto migration tool.
			if err := client.Schema.Create(context.Background(), schema.WithForeignKeys(false)); err != nil {
				return fmt.Errorf("failed creating schema resources: %v", err)
			}
			return nil
		},
		OnStop: func(context.Context) error {
			return client.Close()
		},
	})

	return client, nil
}

func NewRedisClient(lc fx.Lifecycle, _ config.Type) (*redis.Client, error) {
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
