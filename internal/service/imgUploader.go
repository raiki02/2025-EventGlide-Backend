package service

import (
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
)

type ImgUploaderHdl interface {
	GetQIQIUToken(*gin.Context) string
}

type ImgUploader struct {
	AccessKey string
	SerectKey string
	Bucket    string
	ImgUrl    string
}

func NewImgUploader() *ImgUploader {
	return &ImgUploader{
		AccessKey: viper.GetString("imgbed.accecssKey"),
		SerectKey: viper.GetString("imgbed.secretKey"),
		Bucket:    viper.GetString("imgbed.bucket"),
		ImgUrl:    viper.GetString("imgbed.imgUrl"),
	}
}

func (iu *ImgUploader) GenQINIUToken(c *gin.Context) string {
	mac := auth.New(iu.AccessKey, iu.SerectKey)
	putPolicy := storage.PutPolicy{
		Scope:   iu.Bucket,
		Expires: 3600,
	}
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
