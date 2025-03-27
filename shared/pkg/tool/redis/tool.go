package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type tool struct {
	client *redis.Client
	ctx    context.Context
}

func New() Tool {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &tool{client: client, ctx: context.Background()}
}

func (t *tool) Get(key string) (string, error) {
	return t.client.Get(t.ctx, key).Result()
}

func (t *tool) SetEX(key string, value interface{}, expiration time.Duration) error {
	return t.client.SetEx(t.ctx, key, value, expiration).Err()
}

func (t *tool) LPush(key string, value ...interface{}) error {
	return t.client.LPush(t.ctx, key, value).Err()
}

func (t *tool) RPush(key string, value ...interface{}) error {
	return t.client.RPush(t.ctx, key, value).Err()
}

func (t *tool) LRange(key string, start int, stop int) []string {
	return t.client.LRange(t.ctx, key, int64(start), int64(stop)).Val()
}

func (t *tool) HSet(key string, value ...interface{}) error {
	return t.client.HSet(t.ctx, key, value).Err()
}

func (t *tool) HGetAll(key string, dest interface{}) error {
	return t.client.HGetAll(t.ctx, key).Scan(dest)
}

func (t *tool) HGetAllMap(key string) (map[string]string, error) {
	return t.client.HGetAll(t.ctx, key).Result()
}

func (t *tool) Del(keys ...string) error {
	return t.client.Del(t.ctx, keys...).Err()
}
