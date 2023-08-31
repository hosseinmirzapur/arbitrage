package services

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var client *redis.Client

func RedisInit() error {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client.Ping(ctx).Err()
}

func RSet(key string, value interface{}) error {
	return client.Set(ctx, key, value, time.Minute*10).Err()
}

func RGet(key string) (interface{}, error) {
	return client.Get(ctx, key).Result()
}

func RDelete(key string) error {
	return client.Del(ctx, key).Err()
}
