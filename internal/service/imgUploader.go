package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
	"mime/multipart"
	"strconv"
)

type ImgUploaderHdl interface {
	//封装了upload和getfile
	ProcessImg(*gin.Context) (string, error)

	Upload(*gin.Context, *multipart.FileHeader) (string, error)
	GetFile(*gin.Context) []*multipart.FileHeader
}

type ImgUploader struct {
	AccessKey string
	SerectKey string
	Bucket    string
	ImgUrl    string
}

func NewImgUploader() ImgUploaderHdl {
	return &ImgUploader{
		AccessKey: viper.GetString("imgbed.accecssKey"),
		SerectKey: viper.GetString("imgbed.secretKey"),
		Bucket:    viper.GetString("imgbed.bucket"),
		ImgUrl:    viper.GetString("imgbed.imgUrl"),
	}
}

func (iu *ImgUploader) ProcessImg(ctx *gin.Context) (string, error) {
	fhs := iu.GetFile(ctx)
	var urls string
	for _, f := range fhs {
		url, err := iu.Upload(ctx, f)
		if err != nil {
			return "", nil
		}
		//eg."abc.com; qwe.com; "
		urls = url + "; "
	}
	return urls, nil
}

// TODO: 不优雅。感觉跑不起来
func (iu *ImgUploader) GetFile(ctx *gin.Context) []*multipart.FileHeader {
	base := "file"
	var fhs []*multipart.FileHeader
	//获取file0 - file8 的用户传入图片（如果存在）
	for i := 0; i < 9; i++ {
		filename := base + strconv.Itoa(i)
		file, err := ctx.FormFile(filename)
		if err != nil {
			return nil
		}
		fhs = append(fhs, file)
	}
	return fhs
}

func (iu *ImgUploader) Upload(ctx *gin.Context, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	putPlicy := storage.PutPolicy{
		Scope: iu.Bucket,
	}
	mac := qbox.NewMac(iu.AccessKey, iu.SerectKey)

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
		return "", err
	}

	url := iu.ImgUrl + ret.Key // 返回上传后的文件访问路径
	return url, nil
}
