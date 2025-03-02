package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/model"
	"testing"
	"time"
)

func getCtr(t *testing.T) *gomock.Controller {
	return gomock.NewController(t)
}

var ctx *gin.Context

func TestUserDao(t *testing.T) {
	ctr := getCtr(t)
	userDao := NewMockUserDaoHdl(ctr)

	userDao.EXPECT().UpdateAvatar(gomock.Any(), "1234566", "http://img.localhost:9090").Return(nil)
	userDao.UpdateAvatar(ctx, "1234566", "http://img.localhost:9090")

	userDao.EXPECT().UpdateUsername(gomock.Any(), "old_name", "new_name").Return(nil)
	userDao.UpdateUsername(ctx, "old_name", "new_name")

	userDao.EXPECT().Create(gomock.Any(),
		&model.User{
			StudentId: "1234566",
			Name:      "1234566",
			Avatar:    "http://img.localhost:9090",
			School:    "华中师范大学",
		}).
		Return(nil)
	userDao.Create(ctx, &model.User{
		StudentId: "1234566",
		Name:      "1234566",
		Avatar:    "http://img.localhost:9090",
		School:    "华中师范大学",
	})

	userDao.EXPECT().CheckUserExist(gomock.Any(), "1234566").Return(true)
	userDao.CheckUserExist(ctx, "1234566")

	userDao.EXPECT().GetUserInfo(gomock.Any(), "1234566").Return(model.User{}, nil)
	userDao.GetUserInfo(ctx, "1234566")
}

func TestPostDao(t *testing.T) {
	ctr := getCtr(t)
	postDao := NewMockPostDaoHdl(ctr)

	postDao.EXPECT().GetAllPost(gomock.Any()).Return([]model.Post{}, nil)
	postDao.GetAllPost(ctx)

	postDao.EXPECT().CreatePost(gomock.Any(), &model.Post{
		Title:     "test_title",
		Content:   "test_content",
		CreatorID: "1234566",
	}).Return(nil)
	postDao.CreatePost(ctx, &model.Post{
		Title:     "test_title",
		Content:   "test_content",
		CreatorID: "1234566",
	})

	postDao.EXPECT().FindPostByName(gomock.Any(), "test_title").Return([]model.Post{}, nil)
	postDao.FindPostByName(ctx, "test_title")

	postDao.EXPECT().DeletePost(gomock.Any(), &model.Post{
		Bid:       "1234566",
		CreatorID: "1234566",
	}).Return(nil)
	postDao.DeletePost(ctx, &model.Post{
		Bid:       "1234566",
		CreatorID: "1234566",
	})

	postDao.EXPECT().FindPostByUser(gomock.Any(), "1234566", "test_title").Return([]model.Post{}, nil)
	postDao.FindPostByUser(ctx, "1234566", "test_title")

	postDao.EXPECT().CreateDraft(gomock.Any(), &model.PostDraft{
		Title:     "test_title_draft",
		Content:   "test_content_draft",
		CreatorID: "1234566",
	}).Return(nil)

	postDao.CreateDraft(ctx, &model.PostDraft{
		Title:     "test_title_draft",
		Content:   "test_content_draft",
		CreatorID: "1234566",
	})

	postDao.EXPECT().LoadDraft(gomock.Any(), "1234566", "1234566").Return(&model.PostDraft{}, nil)
	postDao.LoadDraft(ctx, "1234566", "1234566")

}

func TestActDao(t *testing.T) {
	ctr := getCtr(t)
	actDao := NewMockActDaoHdl(ctr)

	actDao.EXPECT().CreateAct(gomock.Any(), &model.Activity{
		CreatedAt:      time.Now(),
		CreatorId:      "1234566",
		Bid:            "1-1-1-1",
		Type:           "test_type",
		Host:           "test_host",
		Location:       "test_location",
		IfRegister:     "yes",
		RegisterMethod: "test_register_method",
		StartTime:      "2021-09-01 00:00:00",
		EndTime:        "2021-09-01 00:00:00",
		Name:           "test_name",
		Capacity:       10,
		ImgUrls:        "http://img.localhost:9090",
		Description:    "test_description",
	}).Return(nil)
	actDao.CreateAct(ctx, &model.Activity{
		CreatedAt:      time.Now(),
		CreatorId:      "1234566",
		Bid:            "1-1-1-1",
		Type:           "test_type",
		Host:           "test_host",
		Location:       "test_location",
		IfRegister:     "yes",
		RegisterMethod: "test_register_method",
		StartTime:      "2021-09-01 00:00:00",
		EndTime:        "2021-09-01 00:00:00",
		Name:           "test_name",
		Capacity:       10,
		ImgUrls:        "http://img.localhost:9090",
		Description:    "test_description",
	})

	actDao.EXPECT().CreateDraft(gomock.Any(), &model.ActivityDraft{
		Bid:            "1-1-1-2",
		CreatorID:      "1234566",
		Type:           "test_type",
		Host:           "test_host",
		Location:       "test_location",
		IfRegister:     "yes",
		RegisterMethod: "test_register_method",
		StartTime:      "2021-09-01 00:00:00",
		EndTime:        "2021-09-01 00:00:00",
		Name:           "test_name",
		Capacity:       10,
		Description:    "test_description",
	}).Return(nil)
	actDao.CreateDraft(ctx, &model.ActivityDraft{
		Bid:            "1-1-1-2",
		CreatorID:      "1234566",
		Type:           "test_type",
		Host:           "test_host",
		Location:       "test_location",
		IfRegister:     "yes",
		RegisterMethod: "test_register_method",
		StartTime:      "2021-09-01 00:00:00",
		EndTime:        "2021-09-01 00:00:00",
		Name:           "test_name",
		Capacity:       10,
		Description:    "test_description",
	})

	actDao.EXPECT().LoadDraft(gomock.Any(), "1234566", "1-1-1-2").Return(&model.ActivityDraft{}, nil)
	actDao.LoadDraft(ctx, "1234566", "1-1-1-2")

	actDao.EXPECT().FindActBySearches(gomock.Any(), &req.ActSearchReq{
		Type:       []string{"test_type"},
		Host:       []string{"test_host"},
		Location:   []string{"test_location"},
		IfRegister: "yes",
		StartTime:  "2021-09-01 00:00:00",
		EndTime:    "2021-09-01 00:00:00",
	}).Return([]model.Activity{}, nil)
	actDao.FindActBySearches(ctx, &req.ActSearchReq{
		Type:       []string{"test_type"},
		Host:       []string{"test_host"},
		Location:   []string{"test_location"},
		IfRegister: "yes",
		StartTime:  "2021-09-01 00:00:00",
		EndTime:    "2021-09-01 00:00:00",
	})

	actDao.EXPECT().FindActByDate(gomock.Any(), "2021-09-01").Return([]model.Activity{}, nil)
	actDao.FindActByDate(ctx, "2021-09-01")

	actDao.EXPECT().FindActByName(gomock.Any(), "test_name").Return([]model.Activity{}, nil)
	actDao.FindActByName(ctx, "test_name")

	actDao.EXPECT().DeleteAct(gomock.Any(), model.Activity{
		Type:       "test_type",
		Host:       "test_host",
		Location:   "test_location",
		IfRegister: "yes",
		Capacity:   10,
	}).Return(nil)
	actDao.DeleteAct(ctx, model.Activity{
		Type:       "test_type",
		Host:       "test_host",
		Location:   "test_location",
		IfRegister: "yes",
		Capacity:   10,
	})

	actDao.EXPECT().CheckExist(gomock.Any(), &model.Activity{
		Type:       "test_type",
		Host:       "test_host",
		Location:   "test_location",
		IfRegister: "yes",
		Capacity:   10,
	}).Return(true)
	actDao.CheckExist(ctx, &model.Activity{
		Type:       "test_type",
		Host:       "test_host",
		Location:   "test_location",
		IfRegister: "yes",
		Capacity:   10,
	})

}

func TestCommentDao(t *testing.T) {

}

func TestNumberDao(t *testing.T) {

}
