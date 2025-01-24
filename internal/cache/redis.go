package cache

import (
	"context"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/tools"
	"github.com/redis/go-redis/v9"
	"time"
)

// 先找缓存，缓存没有再找数据库
type CacheHdl interface {
	Get(ctx context.Context, key string) ([]model.Activity, error)
	Set(ctx context.Context, key string, val []model.Activity) error
	Del(ctx context.Context, key string) error
}

// 操作顺序 router -> controller -> cache -> dao
type Cache struct {
	rdb *redis.Client
	//操作评论的点赞和评论
	cd dao.CommentDAOHdl
}

func NewCache(rdb *redis.Client, cd dao.CommentDAOHdl) CacheHdl {
	return &Cache{rdb: rdb, cd: cd}
}

func (c *Cache) Get(ctx context.Context, key string) ([]model.Activity, error) {
	key = "activity:" + key
	b, err := c.rdb.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return jsonTOstrs(b), nil
}

func (c *Cache) Set(ctx context.Context, key string, val []model.Activity) error {
	key = "activity:" + key
	return c.rdb.Set(ctx, key, strsTOjson(val), 48*time.Hour).Err()
}

func (c *Cache) Del(ctx context.Context, key string) error {
	return c.rdb.Del(ctx, key).Err()
}

// redis的val不能存结构体，转换成json格式储存
func strsTOjson(as []model.Activity) []byte {
	return []byte(tools.Marshal(as))
}

func jsonTOstrs(b []byte) []model.Activity {
	res := tools.Unmarshal(b, []model.Activity{})
	as, ok := res.([]model.Activity)
	if !ok {
		return nil
	}
	return as
}
