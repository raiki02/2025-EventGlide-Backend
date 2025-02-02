package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"net/http"
)

var AccessKey = "tgq04Gd9MG-6z9KEM-U9eVTl0zetB_bbIDxTBNQ4"
var SerectKey = "ijalOIFxBETP6NUjSdM3PEKbN-35OWcisSLOkXHf"
var Bucket = "raiki-test"
var ImgUrl = "http://img.inside-me.top"

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("view/index.html")

	// index页面显示
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 表单提交
	r.POST("/uploadfile", func(c *gin.Context) {

		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 10010,
				"msg":  err.Error(),
			})
			return
		}

		// 上传到七牛云
		code, url := UploadToQiNiu(f)

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  "OK",
			"url":  url,
		})

	})
	// 运行，监听127.0.0.1:8080
	r.Run()
}

func UploadToQiNiu(file *multipart.FileHeader) (int, string) {
	src, err := file.Open()
	if err != nil {
		return 10011, err.Error()
	}
	defer src.Close()

	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SerectKey)

	// 获取上传凭证
	upToken := putPlicy.UploadToken(mac)

	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 华南区
		UseCdnDomains: false,
		UseHTTPS:      false, // 非https
	}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}        // 上传后返回的结果
	putExtra := storage.PutExtra{} // 额外参数

	// 上传 自定义key，可以指定上传目录及文件名和后缀，
	key := "image/" + file.Filename // 上传路径，如果当前目录中已存在相同文件，则返回上传失败错误
	err = formUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)

	// 以默认key方式上传
	// err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, src, fileSize, &putExtra)

	// 自定义key，上传指定路径的文件
	// localFilePath = "./aa.jpg"
	// err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFilePath, &putExtra)

	// 默认key，上传指定路径的文件
	// localFilePath = "./aa.jpg"
	// err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFilePath, &putExtra)

	if err != nil {
		code := 501
		return code, err.Error()
	}

	url := ImgUrl + ret.Key // 返回上传后的文件访问路径
	return 0, url
}
