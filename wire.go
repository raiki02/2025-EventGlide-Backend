//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/raiki02/EG/internal/cache"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/ioc"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/raiki02/EG/internal/mq"
	"github.com/raiki02/EG/internal/router"
	"github.com/raiki02/EG/internal/server"
	"github.com/raiki02/EG/internal/service"
)

func InitApp(e *gin.Engine) *server.Server {
	wire.Build(
		ioc.InitDB,
		ioc.InitRedis,
		ioc.Newlogger,
		//ioc.NewKafkaClient,
		//ioc.NewProducer,
		//ioc.NewConsumer,
		//ioc.NewKafka,

		cache.NewCache,
		mq.NewMQ,

		dao.NewFeedDao,
		dao.NewInteractionDao,
		dao.NewActDao,
		dao.NewUserDao,
		dao.NewPostDao,
		dao.NewCommentDao,
		//dao.NewNumberDao,

		service.NewFeedService,
		service.NewInteractionService,
		service.NewImgUploader,
		service.NewPostService,
		service.NewUserService,
		service.NewCCNUService,
		service.NewCommentService,
		service.NewActivityService,

		middleware.NewJwt,
		middleware.NewCors,

		controller.NewFeedController,
		controller.NewInteractionController,
		controller.NewActController,
		controller.NewPostController,
		controller.NewUserController,
		controller.NewCommentController,

		router.NewInteractionRouter,
		router.NewActRouter,
		router.NewCommentRouter,
		router.NewPostRouter,
		router.NewUserRouter,
		router.NewFeedRouter,
		router.NewRouter,

		server.NewServer,
	)
	return &server.Server{}
}
