package ioc

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitRedis() *redis.Client {
	addr := viper.GetString("redis.addr")
	pw := viper.GetString("redis.password")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
	})
	return rdb
}
