package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type CacheHdl interface {
	Get(context.Context, string) ([]interface{}, error)
	Set(context.Context, string, []interface{}) error
	Del(context.Context, string) error
}

type Cache struct {
	rdb *redis.Client
}

func NewCache(rdb *redis.Client) *Cache {
	return &Cache{
		rdb: rdb,
	}
}
