package locker

import (
	"context"
	"time"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
)

type Locker struct {
	lock *redislock.Client
}

func NewLocker(client *redis.Client) *Locker {
	return &Locker{lock: redislock.New(client)}
}

func (l *Locker) Acquire(ctx context.Context, key string, ttl time.Duration) (*redislock.Lock, error) {
	return l.lock.Obtain(ctx, key, ttl, nil)
}
