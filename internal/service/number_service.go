package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/redis/go-redis/v9"
)

type NumberServiceHdl interface {
	SendLikesNum(*gin.Context, *req.NumberReq) error
	SendCommentsNum(*gin.Context, *req.NumberReq) error
	HandleNum() error
}

type NumberService struct {
	rdb *redis.Client
}

func NewNumberService(rdb *redis.Client) *NumberService {
	return &NumberService{
		rdb: rdb,
	}
}

func (ns *NumberService) SendLikesNum(c *gin.Context, nq *req.NumberReq) error {
	ns.rdb.Incr(c.Request.Context(), nq.Topic)
	return nil
}

func (ns *NumberService) SendCommentsNum(c *gin.Context, nq *req.NumberReq) error {
	ns.rdb.Incr(c.Request.Context(), nq.Topic)
	return nil
}
