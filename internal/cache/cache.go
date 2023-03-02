package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"time"

	"github.com/Salam4nder/inventory/internal/config"
	"github.com/Salam4nder/inventory/internal/persistence"

	"github.com/go-redis/redis/v8"
)

// Service is an abstract interface for a caching service.
type Service interface {
	Get(ctx context.Context, uuid string) (
		persistence.Item, error)
	Set(ctx context.Context, key string,
		item persistence.Item, expiration time.Duration) error
	Delete(ctx context.Context, uuid string) error
	Ping(ctx context.Context) error
}

// Redis is an implementation of the cache.Service interface.
type Redis struct {
	Client *redis.Client
}

// New returns a new instance of Redis.
func New(cfg config.Cache) (*Redis, error) {
	redis := &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     cfg.Addr(),
			Password: cfg.Password,
			DB:       0,
		}),
	}

	if err := redis.Client.Ping(
		context.Background()).Err(); err != nil {
		return nil, err
	}

	return redis, nil
}

// Ping checks if Redis is available.
func (r *Redis) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}

// Get checks if a persistence.Item with the given uuid
// exists in the cache. If the key does not exist or if
// there is an error, an empty item and an error are returned.
func (r *Redis) Get(
	ctx context.Context, uuid string) (persistence.Item, error) {
	cmd := r.Client.Get(ctx, uuid)

	// Bytes() will not attempt to convert the command to
	// bytes if there was an error with the GET command,
	// e.g. the key does not exist, or connection error.
	cmdBytes, err := cmd.Bytes()
	if err != nil {
		return persistence.Item{}, err
	}

	bReader := bytes.NewReader(cmdBytes)

	var item persistence.Item

	if err := gob.NewDecoder(bReader).Decode(&item); err != nil {
		return persistence.Item{}, err
	}

	return item, nil
}

// Set caches the given item with the given uuid as the key.
// If the key already exists, it will be overwritten.
// If there is an error, it will be returned.
func (r *Redis) Set(
	ctx context.Context,
	key string,
	item persistence.Item,
	expiration time.Duration) error {
	var buffer bytes.Buffer

	if err := gob.NewEncoder(&buffer).Encode(item); err != nil {
		return err
	}

	return r.Client.Set(
		ctx, key, buffer.Bytes(), expiration).Err()
}

// Delete removes the item with the given uuid from the cache.
// If the key does not exist, an error is returned.
func (r *Redis) Delete(ctx context.Context, uuid string) error {
	res := r.Client.Del(ctx, uuid).Val()
	if res == 0 {
		return errors.New("error deleting key")
	}

	return nil
}
