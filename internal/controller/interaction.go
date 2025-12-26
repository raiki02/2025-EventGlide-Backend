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

type InteractionController struct {
	is *service.InteractionService
	l  *zap.Logger
}

func NewInteractionController(is *service.InteractionService, l *zap.Logger) *InteractionController {
	return &InteractionController{
		is: is,
		l:  l.Named("interaction/controller"),
	}
}

// @Tags Interaction
// @Summary 点赞
// @Accept json
// @Param Authorization header string true "token"
// @Param interaction body req.InteractionReq true "互动"
// @Success 200 {object} resp.Resp
// @Router /interaction/like [post]
func (ic *InteractionController) Like(ctx *gin.Context, req_ req.InteractionReq, claims jwt.RegisteredClaims) (resp.Resp, error) {
	if err := ic.is.Like(ctx, &req_, claims.Subject); err != nil {
		return ginx.ReturnError(err)
	}

	return ginx.ReturnSuccess(nil)
}

// @Tags Interaction
// @Summary 取消点赞
// @Accept json
// @Param Authorization header string true "token"
// @Param interaction body req.InteractionReq true "互动"
// @Success 200 {object} resp.Resp
// @Router /interaction/dislike [post]
func (ic *InteractionController) Dislike(ctx *gin.Context, req_ req.InteractionReq, claims jwt.RegisteredClaims) (resp.Resp, error) {
	if err := ic.is.Dislike(ctx, &req_, claims.Subject); err != nil {
		return ginx.ReturnError(err)
	}

	return ginx.ReturnSuccess(nil)
}

// @Tags Interaction
// @Summary 收藏
// @Accept json
// @Param Authorization header string true "token"
// @Param interaction body req.InteractionReq true "互动"
// @Success 200 {object} resp.Resp
// @Router /interaction/collect [post]
func (ic *InteractionController) Collect(ctx *gin.Context, req_ req.InteractionReq, claims jwt.RegisteredClaims) (resp.Resp, error) {
	if err := ic.is.Collect(ctx, &req_, claims.Subject); err != nil {
		return ginx.ReturnError(err)
	}

	return ginx.ReturnSuccess(nil)
}

// @Tags Interaction
// @Summary 取消收藏
// @Accept json
// @Param Authorization header string true "token"
// @Param interaction body req.InteractionReq true "互动"
// @Success 200 {object} resp.Resp
// @Router /interaction/discollect [post]
func (ic *InteractionController) Discollect(ctx *gin.Context, req_ req.InteractionReq, claims jwt.RegisteredClaims) (resp.Resp, error) {
	if err := ic.is.Discollect(ctx, &req_, claims.Subject); err != nil {
		return ginx.ReturnError(err)
	}

	return ginx.ReturnSuccess(nil)
}

// @Tags Interaction
// @Summary 作为活动填表人批准发表此活动
// @Accept json
// @Param Authorization header string true "token"
// @Param interaction body req.InteractionReq true "互动"
// @Success 200 {object} resp.Resp
// @Router /interaction/approve [post]
func (ic *InteractionController) Approve(ctx *gin.Context, req_ req.InteractionReq, claims jwt.RegisteredClaims) (resp.Resp, error) {
	if err := ic.is.Approve(ctx, &req_, claims.Subject); err != nil {
		return ginx.ReturnError(err)
	}

	return ginx.ReturnSuccess(nil)
}

// @Tags Interaction
// @Summary 作为活动填表人拒绝发表此活动
// @Accept json
// @Param Authorization header string true "token"
// @Param interaction body req.InteractionReq true "互动"
// @Success 200 {object} resp.Resp
// @Router /interaction/reject [post]
func (ic *InteractionController) Reject(ctx *gin.Context, req_ req.InteractionReq, claims jwt.RegisteredClaims) (resp.Resp, error) {
	if err := ic.is.Reject(ctx, &req_, claims.Subject); err != nil {
		return ginx.ReturnError(err)
	}

	return ginx.ReturnSuccess(nil)
}
