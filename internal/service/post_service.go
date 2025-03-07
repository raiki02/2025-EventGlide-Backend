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
	GetAllPost(*gin.Context) ([]resp.ListPostsResp, error)
	CreatePost(*gin.Context, *req.CreatePostReq) (resp.CreatePostResp, error)
	FindPostByName(*gin.Context, string) ([]resp.ListPostsResp, error)
	DeletePost(*gin.Context, *model.Post) error
	CreateDraft(*gin.Context, *req.CreatePostReq) (resp.CreatePostResp, error)
	LoadDraft(*gin.Context, req.DraftReq) (resp.CreatePostResp, error)
	FindPostByOwnerID(*gin.Context, string) ([]resp.ListPostsResp, error)
}

type PostService struct {
	pdh *dao.PostDao
	ud  *dao.UserDao
}

func NewPostService(pdh *dao.PostDao, ud *dao.UserDao) *PostService {
	return &PostService{
		pdh: pdh,
		ud:  ud,
	}
}

func (ps *PostService) GetAllPost(c *gin.Context) ([]resp.ListPostsResp, error) {
	posts, err := ps.pdh.GetAllPost(c)
	if err != nil {
		return nil, err
	}
	res := ps.ToListResp(c, posts)
	return res, nil
}

func (ps *PostService) CreatePost(c *gin.Context, r *req.CreatePostReq) (resp.CreatePostResp, error) {

	post := toPost(r)

	err := ps.pdh.CreatePost(c, post)
	if err != nil {
		return resp.CreatePostResp{}, err
	}

	return ps.toCreateResp(c, post), nil
}

func (ps *PostService) FindPostByName(c *gin.Context, name string) ([]resp.ListPostsResp, error) {
	posts, err := ps.pdh.FindPostByName(c, name)
	if err != nil {
		return nil, err
	}
	res := ps.ToListResp(c, posts)
	return res, nil
}
func (ps *PostService) DeletePost(c *gin.Context, post *model.Post) error {
	err := ps.pdh.DeletePost(c, post)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PostService) CreateDraft(c *gin.Context, r *req.CreatePostReq) (resp.CreatePostResp, error) {
	draft := toDraft(r)
	err := ps.pdh.CreateDraft(c, draft)
	if err != nil {
		return resp.CreatePostResp{}, err
	}
	return ps.toCreateResp(c, draft), nil
}

func (ps *PostService) LoadDraft(c *gin.Context, req req.DraftReq) (resp.CreatePostResp, error) {
	draft, err := ps.pdh.LoadDraft(c, req.Bid, req.Sid)
	if err != nil {
		return resp.CreatePostResp{}, err
	}
	return ps.toCreateResp(c, draft), nil
}

func (ps *PostService) FindPostByOwnerID(c *gin.Context, id string) ([]resp.ListPostsResp, error) {
	posts, err := ps.pdh.FindPostByOwnerID(c, id)
	if err != nil {
		return nil, err
	}
	res := ps.ToListResp(c, posts)
	return res, nil
}

func (ps *PostService) ToListResp(c *gin.Context, posts []model.Post) []resp.ListPostsResp {
	var res []resp.ListPostsResp
	for _, post := range posts {
		res = append(res, ps.toResp(c, post))
	}
	return res
}

func (ps *PostService) toResp(c *gin.Context, post model.Post) resp.ListPostsResp {
	user := ps.ud.FindUserByID(c, post.StudentID)
	var res resp.ListPostsResp
	res.UserInfo.School = user.School
	res.UserInfo.Username = user.Name
	res.UserInfo.Avatar = user.Avatar
	res.UserInfo.StudentID = user.StudentID

	res.Title = post.Title
	res.Introduce = post.Introduce
	res.ShowImg = tools.StringToSlice(post.ShowImg)
	res.LikeNum = post.LikeNum
	res.CommentNum = post.CommentNum
	res.CollectNum = post.CollectNum
	return res
}

func toPost(r *req.CreatePostReq) *model.Post {
	return &model.Post{
		Bid:       tools.GenUUID(),
		CreatedAt: time.Now(),

		StudentID: r.StudentID,
		Title:     r.Title,
		Introduce: r.Introduce,
		ShowImg:   tools.SliceToString(r.ShowImg),
	}
}

func toDraft(r *req.CreatePostReq) *model.PostDraft {
	return &model.PostDraft{
		Bid:       tools.GenUUID(),
		CreatedAt: time.Now(),

		StudentID: r.StudentID,
		Title:     r.Title,
		Introduce: r.Introduce,
		ShowImg:   tools.SliceToString(r.ShowImg),
	}
}

func (ps *PostService) toCreateResp(c *gin.Context, p any) resp.CreatePostResp {
	switch p.(type) {
	case model.Post:
		post := p.(model.Post)
		var res resp.CreatePostResp
		user := ps.ud.FindUserByID(c, post.StudentID)
		res.UserInfo.School = user.School
		res.UserInfo.Username = user.Name
		res.UserInfo.Avatar = user.Avatar
		res.UserInfo.StudentID = user.StudentID
		res.Title = post.Title
		res.Introduce = post.Introduce
		res.ShowImg = tools.StringToSlice(post.ShowImg)
		return res
	case model.PostDraft:
		draft := p.(model.PostDraft)
		var res resp.CreatePostResp
		user := ps.ud.FindUserByID(c, draft.StudentID)
		res.UserInfo.School = user.School
		res.UserInfo.Username = user.Name
		res.UserInfo.Avatar = user.Avatar
		res.UserInfo.StudentID = user.StudentID
		res.Title = draft.Title
		res.Introduce = draft.Introduce
		res.ShowImg = tools.StringToSlice(draft.ShowImg)
		res.Bid = draft.Bid
		return res

	default:
		return resp.CreatePostResp{}
	}
}
