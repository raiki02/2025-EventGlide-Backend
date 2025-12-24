package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/api/resp"
	"github.com/raiki02/EG/internal/service"
	"github.com/raiki02/EG/pkg/ginx"
	"go.uber.org/zap"
)

type FeedController struct {
	fs *service.FeedService
	l  *zap.Logger
}

func NewFeedController(fs *service.FeedService, l *zap.Logger) *FeedController {
	return &FeedController{
		fs: fs,
		l:  l.Named("feed/controller"),
	}
}

// @Summary 获取用户的消息总数
// @Tags feed
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} resp.Resp{data=resp.BriefFeedResp}
// @Router /feed/total [get]
func (fc *FeedController) GetTotalCnt(ctx *gin.Context, claims jwt.RegisteredClaims) (resp.Resp, error) {
	res, err := fc.fs.GetTotalCnt(ctx, claims.Subject)
	if err != nil {
		return ginx.ReturnError(err)
	}

	return ginx.ReturnSuccess(res)
}

// @Summary 获取feed列表
// @Tags feed
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} resp.Resp{data=resp.FeedResp}
// @Router /feed/list [get]
func (fc *FeedController) GetFeedList(ctx *gin.Context, claims jwt.RegisteredClaims) (resp.Resp, error) {
	res, err := fc.fs.GetFeedList(ctx, claims.Subject)
	if err != nil {
		return ginx.ReturnError(err)
	}

	return ginx.ReturnSuccess(res)
}

// @Summary 获取审核员feed列表
// @Tags feed
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} resp.Resp{data=resp.FeedResp}
// @Router /feed/auditor [get]
func (fc *FeedController) GetAuditorFeedList(ctx *gin.Context, claims jwt.RegisteredClaims) (resp.Resp, error) {
	res, err := fc.fs.GetAuditorFeedList(ctx, claims.Subject)
	if err != nil {
		return ginx.ReturnError(err)
	}

	return ginx.ReturnSuccess(res)
}

// @Summary 读取feed详情, 标记已读
// @Tags feed
// @Produce json
// @Param Authorization header string true "token"
// @Param id query string true "业务ID"
// @Success 200 {object} resp.Resp
// @Router /feed/read/detail [get]
func (fc *FeedController) ReadFeedDetail(ctx *gin.Context, req_ req.ReadFeedDetailReq, claims jwt.RegisteredClaims) (resp.Resp, error) {
	if err := fc.fs.ReadFeedDetail(ctx, claims.Subject, req_.Id); err != nil {
		return ginx.ReturnError(err)
	}

	return ginx.ReturnSuccess(nil)
}

// @Summary 读取全部feed, 标记已读
// @Tags feed
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} resp.Resp
// @Router /feed/read/all [get]
func (fc *FeedController) ReadAllFeed(ctx *gin.Context, claims jwt.RegisteredClaims) (resp.Resp, error) {
	if err := fc.fs.ReadAllFeed(ctx, claims.Subject); err != nil {
		return ginx.ReturnError(err)
	}

	return ginx.ReturnSuccess(nil)
}
