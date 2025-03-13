package dao

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"gorm.io/gorm"
)

type PostDaoHdl interface {
	GetAllPost(ctx context.Context) ([]model.Post, error)
	CreatePost(ctx context.Context, post *model.Post) error
	FindPostByName(ctx context.Context, name string) ([]model.Post, error)
	DeletePost(ctx context.Context, post *model.Post) error
	FindPostByUser(ctx context.Context, sid string, keyword string) ([]model.Post, error)
	CreateDraft(ctx context.Context, draft *model.PostDraft) error
	LoadDraft(ctx context.Context, bid string, sid string) (model.PostDraft, error)
	FindPostByOwnerID(ctx context.Context, id string) ([]model.Post, error)
	Like(c *gin.Context, bid string) error
	Comment(c *gin.Context, bid string) error
	Collect(c *gin.Context, bid string) error
}

type PostDao struct {
	db *gorm.DB
}

func NewPostDao(db *gorm.DB) *PostDao {
	return &PostDao{
		db: db,
	}
}

func (pd *PostDao) GetAllPost(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post
	err := pd.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pd *PostDao) CreatePost(ctx context.Context, post *model.Post) error {
	return pd.db.Create(post).Error
}

func (pd *PostDao) FindPostByName(ctx context.Context, name string) ([]model.Post, error) {
	var posts []model.Post
	err := pd.db.Where("title like ?", fmt.Sprintf("%%%s%%", name)).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pd *PostDao) DeletePost(ctx context.Context, post *model.Post) error {
	var p model.Post
	return pd.db.Where("bid = ? and student_id = ?", post.Bid, post.StudentID).Delete(&p).Error
}

func (pd *PostDao) FindPostByUser(ctx context.Context, sid string, keyword string) ([]model.Post, error) {
	if keyword == "" {
		var posts []model.Post
		err := pd.db.Where("student_id = ?", sid).Find(&posts).Error
		if err != nil {
			return nil, err
		}
		return posts, nil
	} else {
		var posts []model.Post
		err := pd.db.Where("student_id = ? and title like ?", sid, fmt.Sprintf("%%%s%%", keyword)).Find(&posts).Error
		if err != nil {
			return nil, err
		}
		return posts, nil
	}
}

func (pd *PostDao) CreateDraft(ctx context.Context, draft *model.PostDraft) error {
	return pd.db.Create(draft).Error
}

func (pd *PostDao) LoadDraft(ctx context.Context, bid string, sid string) (model.PostDraft, error) {
	var draft model.PostDraft
	err := pd.db.Where("bid = ? and student_id = ?", bid, sid).First(&draft).Error
	if err != nil {
		return model.PostDraft{}, err
	}
	return draft, nil
}

func (pd *PostDao) FindPostByOwnerID(ctx context.Context, id string) ([]model.Post, error) {
	var posts []model.Post
	err := pd.db.Where("student_id = ?", id).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pd *PostDao) Like(c *gin.Context, bid string, t int) error {
	var post model.Post
	switch t {
	case 1:
		return pd.db.Where("bid = ?", bid).First(&post).Update("like_num", post.LikeNum+1).Error
	case 0:
		return pd.db.Where("bid = ?", bid).First(&post).Update("like_num", post.LikeNum-1).Error
	default:
		return fmt.Errorf("type error")
	}
}

func (pd *PostDao) Comment(c *gin.Context, bid string, t int) error {
	var post model.Post
	switch t {
	case 1:
		return pd.db.Where("bid = ?", bid).First(&post).Update("comment_num", post.CommentNum+1).Error
	case 0:
		return pd.db.Where("bid = ?", bid).First(&post).Update("comment_num", post.CommentNum-1).Error
	default:
		return fmt.Errorf("type error")
	}
}
func (pd *PostDao) Collect(c *gin.Context, bid string, t int) error {
	var post model.Post
	switch t {
	case 1:
		return pd.db.Where("bid = ?", bid).First(&post).Update("collect_num", post.CollectNum+1).Error
	case 0:
		return pd.db.Where("bid = ?", bid).First(&post).Update("collect_num", post.CollectNum-1).Error
	default:
		return fmt.Errorf("type error")
	}
}
func (pd *PostDao) FindPostByBid(c *gin.Context, bid string) (model.Post, error) {
	var post model.Post
	err := pd.db.Where("bid = ?", bid).First(&post).Error
	if err != nil {
		return model.Post{}, err
	}
	return post, nil
}
