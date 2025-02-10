package cache

import (
	"context"
	"github.com/raiki02/EG/tools"
	"github.com/redis/go-redis/v9"
	"time"
)

// 先找缓存，缓存没有再找数据库
// TODO: scpoe： 活动，帖子，jwt
type CacheHdl interface {
	Get(context.Context, string) ([]interface{}, error)
	Set(context.Context, string, []interface{}) error
	Del(context.Context, string) error
}

// 操作顺序 router -> controller -> cache -> dao
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

// redis的val不能存结构体，转换成json格式储存
func Tojson(in []interface{}) []byte {
	return []byte(tools.Marshal(in))
}

func jsonTOstrus(b []byte) (out []interface{}) {
	return nil
}
