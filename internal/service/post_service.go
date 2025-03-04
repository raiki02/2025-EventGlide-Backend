package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/api/resp"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/tools"
	"time"
)

type PostServiceHdl interface {
	CreatePost(*gin.Context, *model.Post) error
	DeletePost(*gin.Context, *model.Post) error

	CreateDraft(*gin.Context, *model.PostDraft) error
	LoadDraft(*gin.Context, req.DraftReq) (*model.PostDraft, error)

	FindPostByOwnerID(*gin.Context, string) ([]resp.ListPostsResp, error)
	FindPostByName(*gin.Context, string) ([]resp.ListPostsResp, error)
	GetAllPost(*gin.Context) ([]resp.ListPostsResp, error)
}

type PostService struct {
	pdh *dao.PostDao
	ud  *dao.UserDao
}

func NewPostService(pdh *dao.PostDao) *PostService {
	return &PostService{
		pdh: pdh,
	}
}

func (ps *PostService) GetAllPost(c *gin.Context) ([]resp.ListPostsResp, error) {
	posts, err := ps.pdh.GetAllPost(c)
	if err != nil {
		return nil, err
	}
	res := ps.ToResp(c, posts)
	return res, nil
}

func (ps *PostService) CreatePost(c *gin.Context, post *model.Post) error {
	post.Bid = tools.GenUUID(c)
	err := ps.pdh.CreatePost(c, post)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PostService) FindPostByName(c *gin.Context, name string) ([]resp.ListPostsResp, error) {
	posts, err := ps.pdh.FindPostByName(c, name)
	if err != nil {
		return nil, err
	}
	res := ps.ToResp(c, posts)
	return res, nil
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
	draft.CreatedAt = time.Now()
	err := ps.pdh.CreateDraft(c, draft)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PostService) LoadDraft(c *gin.Context, req req.DraftReq) (model.PostDraft, error) {
	draft, err := ps.pdh.LoadDraft(c, req.Bid, req.Sid)
	if err != nil {
		return model.PostDraft{}, err
	}
	return draft, nil
}

func (ps *PostService) FindPostByOwnerID(c *gin.Context, id string) ([]resp.ListPostsResp, error) {
	posts, err := ps.pdh.FindPostByOwnerID(c, id)
	if err != nil {
		return nil, err
	}
	res := ps.ToResp(c, posts)
	return res, nil
}

func (ps *PostService) ToResp(c *gin.Context, posts []model.Post) []resp.ListPostsResp {
	var res []resp.ListPostsResp
	for n, post := range posts {
		user := ps.ud.FindUserByID(c, post.CreatorID)
		res[n].User.School = user.School
		res[n].User.Username = user.Name
		res[n].User.Avatar = user.Avatar
		res[n].User.Sid = user.StudentId

		res[n].Title = post.Title
		res[n].Content = post.Content
		res[n].ImgUrls = processUrls(c, post.ImgUrls)
		res[n].Likes = post.Likes
		res[n].Comments = post.Comments
	}
	return res
}
