package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/redis/go-redis/v9"
)

type NumberServiceHdl interface {
	AddLikesNum(*gin.Context, *model.Number) error
	CutLikesNum(*gin.Context, *model.Number) error
	AddCommentsNum(*gin.Context, *model.Number) error
}

type NumberService struct {
	nd  *dao.NumberDao
	rdb *redis.Client
	ad  *dao.ActDao
	pd  *dao.PostDao
	cd  *dao.CommentDao
}

func NewNumberService(nd *dao.NumberDao, rdb *redis.Client, ad *dao.ActDao, pd *dao.PostDao, cd *dao.CommentDao) *NumberService {
	return &NumberService{
		nd:  nd,
		rdb: rdb,
		ad:  ad,
		pd:  pd,
		cd:  cd,
	}
}

func (ns *NumberService) AddLikesNum(c *gin.Context, nr *model.Number) error {
	nr.Topic = "like"
	return ns.nd.AddLikesNum(c, nr)
}

func (ns *NumberService) CutLikesNum(c *gin.Context, nr *model.Number) error {
	nr.Topic = "like"
	return ns.nd.CutLikesNum(c, nr.Sid, nr.Bid)
}

func (ns *NumberService) AddCommentsNum(c *gin.Context, nr *model.Number) error {
	nr.Topic = "comment"
	return ns.nd.AddCommentsNum(c, nr)
}

func (ns *NumberService) GetLikesNum(c *gin.Context, sid string, bid string) int {
	return ns.nd.GetLikesNum(c, sid, bid)
}

func (ns *NumberService) GetCommentsNum(c *gin.Context, sid string, bid string) int {
	return ns.nd.GetCommentsNum(c, sid, bid)
}

func (ns *NumberService) UpdateNumbers(c *gin.Context, sid string, bid string) error {
	likes := ns.GetLikesNum(c, sid, bid)
	comments := ns.GetCommentsNum(c, sid, bid)
	ns.cd.UpdateNumbersForAnswers(c, sid, bid, likes)
	ns.cd.UpdateNumbersForComments(c, sid, bid, likes, comments)
	ns.ad.UpdateNumbers(c, sid, bid, likes, comments)
	ns.pd.UpdateNumbers(c, sid, bid, likes, comments)
	return nil
}
