package ioc

import (
	"github.com/google/wire"
	"github.com/raiki02/EG/internal/cache"
	"github.com/raiki02/EG/internal/mq"
)

var Provider = wire.NewSet(
	InitDB,
	InitRedis,
	Newlogger,
	InitListener,
	InitApiKeyGetter,
	mq.NewMQ,
	cache.NewCache,
)
