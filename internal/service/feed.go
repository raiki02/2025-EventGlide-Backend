package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/resp"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/raiki02/EG/internal/mq"
	"github.com/raiki02/EG/tools"
	"go.uber.org/zap"
	"time"
)

type FeedServiceHdl interface {
	GetTotalCnt(ctx *gin.Context, sid string) (resp.BriefFeedResp, error)
	GetFeedList(ctx *gin.Context, sid string) (resp.FeedResp, error)
	SubsribeTopics(ctx *gin.Context)
	GetLikeFeed(ctx *gin.Context, sid string) ([]resp.FeedLikeResp, error)
	GetCollectFeed(ctx *gin.Context, sid string) ([]resp.FeedCollectResp, error)
	GetCommentFeed(ctx *gin.Context, sid string) ([]resp.FeedCommentResp, error)
	GetAtFeed(ctx *gin.Context, sid string) ([]resp.FeedAtResp, error)
}

type FeedService struct {
	fd *dao.FeedDao
	mq *mq.MQ
	ud *dao.UserDao
	l  *zap.Logger
}

func NewFeedService(fd *dao.FeedDao, mq *mq.MQ, ud *dao.UserDao, l *zap.Logger) *FeedService {
	fs := &FeedService{
		fd: fd,
		mq: mq,
		ud: ud,
		l:  l.Named("feed/service"),
	}
	fs.SubsribeTopics()
	return fs
}

func (fs *FeedService) GetTotalCnt(ctx *gin.Context, sid string) (resp.BriefFeedResp, error) {

	ints, err := fs.fd.GetTotalCnt(ctx, sid)
	if err != nil {
		fs.l.Error("Get All Events Failed", zap.Error(err))
		return resp.BriefFeedResp{}, err
	}
	return resp.BriefFeedResp{
		LikeAndCollect: ints[0],
		CommentAndAt:   ints[1],
		Total:          ints[2],
	}, nil

}

func (fs *FeedService) GetFeedList(ctx *gin.Context, sid string) (resp.FeedResp, error) {
	l, err1 := fs.GetLikeFeed(ctx, sid)
	c, err2 := fs.GetCollectFeed(ctx, sid)
	cm, err3 := fs.GetCommentFeed(ctx, sid)
	a, err4 := fs.GetAtFeed(ctx, sid)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fs.l.Error("Get Feed List Failed", zap.Error(err1), zap.Error(err2), zap.Error(err3), zap.Error(err4))
		return resp.FeedResp{}, errors.New("get feed list error")
	}
	return resp.FeedResp{
		Likes:    l,
		Ats:      a,
		Comments: cm,
		Collects: c,
	}, nil
}

func (fs *FeedService) SubsribeTopics() {
	go func() {
		ctx := context.Background()
		lc := fs.mq.Subscribe(ctx, "feed")

		for msg := range lc.Channel() {
			var r struct {
				StudentID string `json:"studentid"`
				TargetId  string `json:"target_id"`
				Object    string `json:"object"`
				Action    string `json:"action"`
			}

			json.Unmarshal([]byte(msg.Payload), &r)
			feed := model.Feed{
				TargetBid: r.TargetId,
				StudentId: r.StudentID,
				Object:    r.Object,
				CreatedAt: time.Now(),
				Status:    "未读",
				Action:    r.Action,
			}
			err := fs.fd.CreateFeed(ctx, &feed)
			if err != nil {
				fs.l.Error("Consume Feed Failed", zap.Error(err))
			}
			fs.l.Info("Consume Feed Success",
				zap.String("bid", feed.TargetBid),
				zap.String("studentid", feed.StudentId),
				zap.String("object", feed.Object),
				zap.String("action", feed.Action),
			)
		}
	}()
}

func (fs *FeedService) GetLikeFeed(ctx *gin.Context, sid string) ([]resp.FeedLikeResp, error) {
	likes, err := fs.fd.GetLikeFeed(ctx, sid)
	if err != nil {
		fs.l.Error("Get Like Feed List Failed", zap.Error(err))
		return nil, err
	}
	var res []resp.FeedLikeResp
	for _, v := range likes {
		user, err := fs.ud.GetUserInfo(ctx, v.StudentId)
		if err != nil {
			fs.l.Error("Get User Info when get like feed Failed", zap.Error(err))
			return nil, err
		}
		res = append(res, resp.FeedLikeResp{
			Userinfo: resp.UserInfo{
				StudentID: user.StudentID,
				Avatar:    user.Avatar,
				Username:  user.Name,
			},
			Message:     processMsg(v, user.Name),
			PubLishedAt: tools.ParseTime(v.CreatedAt),
			TargetBid:   v.TargetBid,
		})
	}
	return res, nil
}

func (fs *FeedService) GetCollectFeed(ctx *gin.Context, sid string) ([]resp.FeedCollectResp, error) {
	collects, err := fs.fd.GetCollectFeed(ctx, sid)
	if err != nil {
		fs.l.Error("Get Collect Feed List Failed", zap.Error(err))
		return nil, err
	}
	var res []resp.FeedCollectResp
	for _, v := range collects {
		user, err := fs.ud.GetUserInfo(ctx, v.StudentId)
		if err != nil {
			fs.l.Error("Get User Info when get collect feed Failed", zap.Error(err))
			return nil, err
		}
		res = append(res, resp.FeedCollectResp{
			Userinfo: resp.UserInfo{
				StudentID: user.StudentID,
				Avatar:    user.Avatar,
				Username:  user.Name,
			},
			Message:     processMsg(v, user.Name),
			PubLishedAt: tools.ParseTime(v.CreatedAt),
			TargetBid:   v.TargetBid,
		})
	}
	return res, nil
}

func (fs *FeedService) GetCommentFeed(ctx *gin.Context, sid string) ([]resp.FeedCommentResp, error) {
	comments, err := fs.fd.GetCommentFeed(ctx, sid)
	if err != nil {
		fs.l.Error("Get Comment Feed List Failed", zap.Error(err))
		return nil, err
	}
	var res []resp.FeedCommentResp
	for _, v := range comments {
		user, err := fs.ud.GetUserInfo(ctx, v.StudentId)
		if err != nil {
			fs.l.Error("Get User Info when get comment feed Failed", zap.Error(err))
			return nil, err
		}
		res = append(res, resp.FeedCommentResp{
			Userinfo: resp.UserInfo{
				StudentID: user.StudentID,
				Avatar:    user.Avatar,
				Username:  user.Name,
			},
			Message:     processMsg(v, user.Name),
			PubLishedAt: tools.ParseTime(v.CreatedAt),
			TargetBid:   v.TargetBid,
		})
	}
	return res, nil
}

func (fs *FeedService) GetAtFeed(ctx *gin.Context, sid string) ([]resp.FeedAtResp, error) {
	ats, err := fs.fd.GetAtFeed(ctx, sid)
	if err != nil {
		fs.l.Error("Get At Feed List Failed", zap.Error(err))
		return nil, err
	}
	var res []resp.FeedAtResp
	for _, v := range ats {
		user, err := fs.ud.GetUserInfo(ctx, v.StudentId)
		if err != nil {
			fs.l.Error("Get User Info when get at feed Failed", zap.Error(err))
			return nil, err
		}
		res = append(res, resp.FeedAtResp{
			Userinfo: resp.UserInfo{
				StudentID: user.StudentID,
				Avatar:    user.Avatar,
				Username:  user.Name,
			},
			Message:     processMsg(v, user.Name),
			PubLishedAt: tools.ParseTime(v.CreatedAt),
			TargetBid:   v.TargetBid,
		})
	}
	return res, nil
}

func processMsg(f *model.Feed, name string) string {
	switch f.Action {
	case "like":
		switch f.Object {
		case "post":
			return name + "赞了你的帖子"
		case "comment":
			return name + "赞了你的评论"
		case "activity":
			return name + "赞了你的活动"
		}
	case "collect":
		switch f.Object {
		case "post":
			return name + "收藏了你的帖子"
		case "activity":
			return name + "收藏了你的活动"
		}
	case "comment":
		switch f.Object {
		case "post":
			return name + "评论了你的帖子"
		case "comment":
			return name + "评论了你的评论"
		case "activity":
			return name + "评论了你的活动"
		}
	case "at":
		switch f.Object {
		case "comment":
			return name + "在评论中@了你"
		}
	}
	return "消息加载中......"
}
