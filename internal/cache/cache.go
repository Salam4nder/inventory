package cache

import (
	"context"
	"time"

	"github.com/Salam4nder/inventory/config"

	"github.com/go-redis/redis/v8"
)

// Cache is an abstract interface for a cache.
type Cache interface {
	Delete(key string) int64
	Set(
		key string, value interface{}, expiryTime time.Duration) error
	Get(Key string) string
}

// Redis is an implementation of Cache interface.
type Redis struct {
	Client *redis.Client
}

// New returns a new instance of Redis.
func New(cfg config.Cache) (*Redis, error) {
	redis := &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr: cfg.Host,
		}),
	}

	if err := redis.Client.Ping(
		context.Background()).Err(); err != nil {
		return nil, err
	}

	return redis, nil
}

// Get returns the value of the key.
func (c *Redis) Get(ctx context.Context, Key string) string {
	return c.Client.Get(ctx, Key).Val()
}

// Set sets the value of the key.
func (c *Redis) Set(
	ctx context.Context, key string, value interface{},
	expiryTime time.Duration) error {
	return c.Client.Set(ctx, key, value, expiryTime).Err()
}

// Delete deletes the key.
func (c *Redis) Delete(ctx context.Context, key string) int64 {
	return c.Client.Del(ctx, key).Val()
}
