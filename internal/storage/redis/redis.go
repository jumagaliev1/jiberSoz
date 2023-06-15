package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func New(uri string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: "",
		DB:       0,
	})

	return &RedisClient{
		Client: client,
	}
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	err := r.Client.Set(ctx, key, value, 0).Err()

	return err
}

func (r *RedisClient) Get(ctx context.Context, key string) (interface{}, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}
