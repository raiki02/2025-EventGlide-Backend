package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"net/http"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("index.html")

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

		code, url := UploadToQiNiu(f)

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  "OK",
			"url":  url,
		})

	})
	r.Run()
}

func UploadToQiNiu(file *multipart.FileHeader) (int, string) {
	var AccessKey = ""
	var SerectKey = ""
	var Bucket = ""
	var ImgUrl = ""

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
		UseHTTPS:      false,
	}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}        // 上传后返回的结果
	putExtra := storage.PutExtra{} // 额外参数

	// 上传 自定义key，可以指定上传目录及文件名和后缀，
	key := "EventGlide/" + file.Filename // 上传路径，如果当前目录中已存在相同文件，则返回上传失败错误
	err = formUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)

	if err != nil {
		code := 501
		return code, err.Error()
	}

	url := ImgUrl + ret.Key // 返回上传后的文件访问路径
	return 0, url
}
