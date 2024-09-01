package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisAdapter struct {
	client *redis.Client
}

func NewRedisAdapter() *RedisAdapter {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	return &RedisAdapter{client: rdb}
}

func (r *RedisAdapter) Set(key string, value string) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *RedisAdapter) Get(key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisAdapter) Delete(key string) error {
	return r.client.Del(ctx, key).Err()
}
