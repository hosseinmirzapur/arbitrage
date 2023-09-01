package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var client *redis.Client

func RedisInit() error {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		return errors.New("REDIS_ADDR not set")
	}

	pass := os.Getenv("REDIS_PASS")
	if pass == "" {
		return errors.New("REDIS_PASS not set")
	}
	port := os.Getenv("REDIS_PORT")
	if port == "" {
		return errors.New("REDIS_PORT not set")
	}
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", addr, port),
		Password: pass,
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
