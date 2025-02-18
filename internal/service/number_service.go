package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/redis/go-redis/v9"
)

type NumberServiceHdl interface {
	AddLikesNum(*gin.Context, *req.NumberReq) error
	CutLikesNum(*gin.Context, *req.NumberReq) error
	AddCommentsNum(*gin.Context, *req.NumberReq) error
}

type NumberService struct {
	rdb *redis.Client
}

// kafka -> redis -> mysql
func NewNumberService(rdb *redis.Client) *NumberService {
	return &NumberService{
		rdb: rdb,
	}
}

func (ns *NumberService) AddLikesNum(c *gin.Context, nr *req.NumberReq) error {

}

func (ns *NumberService) CutLikesNum(c *gin.Context, nr *req.NumberReq) error {

}

func (ns *NumberService) AddCommentsNum(c *gin.Context, nr *req.NumberReq) error {

}

// consume kafka message -> redis -> mysql
