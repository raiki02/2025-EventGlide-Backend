package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
)

type PostServiceHdl interface {
	GetAllPost(*gin.Context) ([]model.Post, error)
	CreatePost(*gin.Context, *model.Post) error
	FindPostByName(*gin.Context, string) ([]model.Post, error)
	DeletePost(*gin.Context, *model.Post) error
}

type PostService struct {
	pdh dao.PostDaoHdl
	iuh ImgUploaderHdl
}

func NewPostService(pdh dao.PostDaoHdl, iuh ImgUploaderHdl) PostServiceHdl {
	return &PostService{
		pdh: pdh,
		iuh: iuh,
	}
}

func (ps *PostService) GetAllPost(c *gin.Context) ([]model.Post, error) {
	posts, err := ps.pdh.GetAllPost(c)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (ps *PostService) CreatePost(c *gin.Context, post *model.Post) error {
	urls, err := ps.iuh.ProcessImg(c)
	if err != nil {
		return errors.New("img upload failed")
	}
	post.ImgUrls = urls
	err = ps.pdh.CreatePost(c, post)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PostService) FindPostByName(c *gin.Context, name string) ([]model.Post, error) {
	posts, err := ps.pdh.FindPostByName(c, name)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
func (ps *PostService) DeletePost(c *gin.Context, post *model.Post) error {
	err := ps.pdh.DeletePost(c, post)
	if err != nil {
		return err
	}
	return nil
}
