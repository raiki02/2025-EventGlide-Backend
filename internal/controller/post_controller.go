package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/internal/service"
	"github.com/raiki02/EG/tools"
)

type PostControllerHdl interface {
	GetAllPost() gin.HandlerFunc
	CreatePost() gin.HandlerFunc
	FindPostByName() gin.HandlerFunc
	DeletePost() gin.HandlerFunc
}

type PostController struct {
	ps service.PostServiceHdl
}

func NewPostController(ps service.PostServiceHdl) PostControllerHdl {
	return &PostController{
		ps: ps,
	}
}

func (pc *PostController) GetAllPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		posts, err := pc.ps.GetAllPost(c)
		if err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		tools.ReturnMSG(c, "success", posts)
	}
}

func (pc *PostController) CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post model.Post
		err := c.ShouldBindJSON(&post)
		if err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		err = pc.ps.CreatePost(c, &post)
		if err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		tools.ReturnMSG(c, "success", post)
	}
}

func (pc *PostController) FindPostByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		posts, err := pc.ps.FindPostByName(c, name)
		if err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		tools.ReturnMSG(c, "success", posts)
	}
}

func (pc *PostController) DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post model.Post
		err := c.ShouldBindJSON(&post)
		if err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		err = pc.ps.DeletePost(c, &post)
		if err != nil {
			tools.ReturnMSG(c, err.Error(), nil)
			return
		}
		tools.ReturnMSG(c, "success", nil)
	}
}
