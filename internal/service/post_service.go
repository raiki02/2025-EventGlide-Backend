package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/tools"
)

type PostServiceHdl interface {
	GetAllPost(*gin.Context) ([]model.Post, error)
	CreatePost(*gin.Context, *model.Post) error
	FindPostByName(*gin.Context, string) ([]model.Post, error)
	DeletePost(*gin.Context, *model.Post) error
	CreateDraft(*gin.Context, *model.PostDraft) error
	LoadDraft(*gin.Context, req.DraftReq) (*model.PostDraft, error)
}

type PostService struct {
	pdh *dao.PostDao
	iuh *ImgUploader
}

func NewPostService(pdh *dao.PostDao, iuh *ImgUploader) *PostService {
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

func (ps *PostService) CreateDraft(c *gin.Context, draft *model.PostDraft) error {
	draft.Bid = tools.GenUUID(c)
	err := ps.pdh.CreateDraft(c, draft)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PostService) LoadDraft(c *gin.Context, req req.DraftReq) (*model.PostDraft, error) {
	draft, err := ps.pdh.LoadDraft(c, req.Bid, req.Sid)
	if err != nil {
		return nil, err
	}
	return draft, nil
}
