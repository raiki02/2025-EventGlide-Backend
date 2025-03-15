package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/ioc"
)

type NumberServiceHdl interface {
	Send(*gin.Context, req.NumberSendReq) error
	Recv() error
	Delete(*gin.Context, req.NumberDelReq) error
	Search(*gin.Context, req.NumberSearchReq) error
}

type NumberService struct {
	nd *dao.NumberDao
	k  *ioc.Kafka
}

func NewNumberService(nd *dao.NumberDao, k *ioc.Kafka) *NumberService {
	ns := &NumberService{
		nd: nd,
		k:  k,
	}
	return ns
}
