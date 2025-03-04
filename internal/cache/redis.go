package cache

import (
	"context"
	"github.com/raiki02/EG/tools"
	"github.com/redis/go-redis/v9"
	"time"
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

func (c *Cache) Get(ctx context.Context, key string) ([]interface{}, error) {
	b, err := c.rdb.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return jsonTOstrus(b), nil
}

func (c *Cache) Set(ctx context.Context, key string, val []interface{}) error {
	return c.rdb.Set(ctx, key, Tojson(val), 48*time.Hour).Err()
}

func (c *Cache) Del(ctx context.Context, key string) error {
	return c.rdb.Del(ctx, key).Err()
}

func Tojson(in []interface{}) []byte {
	return []byte(tools.Marshal(in))
}

func jsonTOstrus(b []byte) (out []interface{}) {
	return nil
}
