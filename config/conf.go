package config

import (
	"flag"
	"github.com/spf13/viper"
)

var (
	ConfPath string
)

func Init() error {
	flag.StringVar(&ConfPath, "conf", "config/conf.yaml", "配置文件路径")
	flag.Parse()
	viper.SetConfigFile(ConfPath)

	viper.SetDefault("mysql.dsn", "root:114514@tcp(127.0.0.1:3306)/EventGlide?charset=utf8mb4&parseTime=True&loc=Local")
	viper.SetDefault("mysql.maxIdleConns", 10)
	viper.SetDefault("mysql.maxOpenConns", 100)
	viper.SetDefault("redis.addr", "127.0.0.1:6379")

	return viper.ReadInConfig()
}
