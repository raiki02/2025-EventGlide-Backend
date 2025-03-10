// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/cache"
	"github.com/raiki02/EG/internal/controller"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/ioc"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/raiki02/EG/internal/router"
	"github.com/raiki02/EG/internal/server"
	"github.com/raiki02/EG/internal/service"
)

// Injectors from wire.go:

func InitApp(e *gin.Engine) *server.Server {
	db := ioc.InitDB()
	userDao := dao.NewUserDao(db)
	actDao := dao.NewActDao(db)
	postDao := dao.NewPostDao(db)
	commentDao := dao.NewCommentDao(db)
	client := ioc.InitRedis()
	jwt := middleware.NewJwt(client)
	imgUploader := service.NewImgUploader()
	ccnuService := service.NewCCNUService()
	userService := service.NewUserService(userDao, actDao, postDao, commentDao, jwt, imgUploader, ccnuService)
	userController := controller.NewUserController(e, userService)
	userRouter := router.NewUserRouter(e, userController, jwt)
	cacheCache := cache.NewCache(client)
	activityService := service.NewActivityService(actDao, cacheCache, userDao)
	actController := controller.NewActController(activityService, imgUploader)
	actRouter := router.NewActRouter(e, actController, jwt)
	postService := service.NewPostService(postDao, userDao)
	postController := controller.NewPostController(postService)
	postRouter := router.NewPostRouter(e, postController, jwt)
	commentService := service.NewCommentService(commentDao, userDao)
	commentController := controller.NewCommentController(commentService)
	commentRouter := router.NewCommentRouter(commentController, e, jwt)
	numberDao := dao.NewNumberDao(db)
	saramaClient := ioc.NewKafkaClient()
	syncProducer := ioc.NewProducer(saramaClient)
	consumer := ioc.NewConsumer(saramaClient)
	kafka := ioc.NewKafka(syncProducer, consumer)
	numberService := service.NewNumberService(numberDao, kafka)
	numberController := controller.NewNumberController(numberService)
	numberRouter := router.NewNumberRouter(numberController, e, jwt)
	cors := middleware.NewCors(e)
	routerRouter := router.NewRouter(e, userRouter, actRouter, postRouter, commentRouter, numberRouter, cors)
	serverServer := server.NewServer(routerRouter)
	return serverServer
}
