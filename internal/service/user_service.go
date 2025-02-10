package service

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/raiki02/EG/internal/model"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type UserServiceHdl interface {
	CreateUser(*gin.Context, string) error
	Login(*gin.Context, string, string) (*model.User, string, error)
	Logout(*gin.Context) error
	GetUserInfo(*gin.Context, string) (model.User, error)
	UpdateAvatar(*gin.Context, string) error
	UpdateUsername(*gin.Context, string, string) error
	SearchUserAct(*gin.Context, string, string) ([]model.Activity, error)
	SearchUserPost(*gin.Context, string, string) ([]model.Post, error)
}

type UserService struct {
	udh  dao.UserDAOHdl
	adh  dao.ActDaoHdl
	pdh  dao.PostDaoHdl
	jwth middleware.JwtHdl
	cSvc ccnuService
	iuh  ImgUploaderHdl
}

func NewUserService(udh dao.UserDAOHdl, jwth middleware.JwtHdl, cSvc ccnuService, iuh ImgUploaderHdl) UserServiceHdl {
	return &UserService{
		udh:  udh,
		jwth: jwth,
		cSvc: cSvc,
		iuh:  iuh,
	}
}

func (s *UserService) CreateUser(ctx *gin.Context, sid string) error {
	user := &model.User{
		StudentId: sid,
		Name:      sid,
		Avatar:    model.DefaultAvatar,
		School:    "华中师范大学",
	}
	err := s.udh.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(ctx *gin.Context, studentId string, password string) (*model.User, string, error) {
	client, err := s.cSvc.Login(ctx, studentId, password)
	if err != nil {
		return nil, "", err
	}
	if !client {
		return nil, "", errors.New("登录失败")
	}
	if !s.udh.CheckUserExist(ctx, studentId) {
		err := s.CreateUser(ctx, studentId)
		if err != nil {
			return nil, "", err
		}
	}
	token := s.jwth.GenToken(ctx, studentId)
	err = s.jwth.StoreInRedis(ctx, studentId, token)
	if err != nil {
		return nil, "", err
	}
	user, err := s.udh.GetUserInfo(ctx, studentId)
	return &user, token, nil
}

func (s *UserService) Logout(ctx *gin.Context) error {
	token := ctx.GetHeader("Authorization")
	err := s.jwth.ClearToken(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserInfo(ctx *gin.Context, studentId string) (model.User, error) {
	user, err := s.udh.GetUserInfo(ctx, studentId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserService) UpdateAvatar(ctx *gin.Context, sid string) error {
	imgurl, err := s.iuh.ProcessImg(ctx)
	err = s.udh.UpdateAvatar(ctx, sid, imgurl)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUsername(ctx *gin.Context, studentId string, name string) error {
	err := s.udh.UpdateUsername(ctx, studentId, name)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) SearchUserAct(ctx *gin.Context, studentId string, keyword string) ([]model.Activity, error) {
	acts, err := s.adh.FindActByUser(ctx, studentId, keyword)
	if err != nil {
		return nil, err
	}
	return acts, nil
}

func (s *UserService) SearchUserPost(ctx *gin.Context, studentId string, keyword string) ([]model.Post, error) {
	posts, err := s.pdh.FindPostByUser(ctx, studentId, keyword)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

//---一站式账号登录

type ccnuService struct {
	timeout time.Duration
}

func NewCCNUService(timeout time.Duration) *ccnuService {
	return &ccnuService{
		timeout: timeout,
	}
}

func (c *ccnuService) Login(ctx context.Context, studentId string, password string) (bool, error) {
	var (
		client *http.Client
		err    error
	)
	client, err = c.loginUndergraduateClient(ctx, studentId, password)
	return client != nil, err
}

func (c *ccnuService) client() *http.Client {
	j, _ := cookiejar.New(&cookiejar.Options{})
	return &http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
		Jar:     j,
		Timeout: c.timeout,
	}
}

func (c *ccnuService) loginUndergraduateClient(ctx context.Context, studentId string, password string) (*http.Client, error) {
	params, err := c.makeAccountPreflightRequest()
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Set("username", studentId)
	v.Set("password", password)
	v.Set("lt", params.lt)
	v.Set("execution", params.execution)
	v.Set("_eventId", params._eventId)
	v.Set("submit", params.submit)

	request, err := http.NewRequest("POST", "https://account.ccnu.edu.cn/cas/login;jsessionid="+params.JSESSIONID, strings.NewReader(v.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	request.WithContext(ctx)

	client := c.client()
	resp, err := client.Do(request)
	if err != nil {
		var opErr *net.OpError
		if errors.As(err, &opErr) {
			return nil, errors.New("网络异常")
		}
		return nil, err
	}
	if len(resp.Header.Get("Set-Cookie")) == 0 {
		return nil, errors.New("学号或密码错误")
	}
	return client, nil
}

type accountRequestParams struct {
	lt         string
	execution  string
	_eventId   string
	submit     string
	JSESSIONID string
}

func (c *ccnuService) makeAccountPreflightRequest() (*accountRequestParams, error) {
	var JSESSIONID string
	var lt string
	var execution string
	var _eventId string

	params := &accountRequestParams{}

	// 初始化 http request
	request, err := http.NewRequest("GET", "https://account.ccnu.edu.cn/cas/login", nil)
	if err != nil {
		return params, err
	}

	// 发起请求
	resp, err := c.client().Do(request)
	if err != nil {
		return params, err
	}

	// 读取 Body
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return params, err
	}

	// 获取 Cookie 中的 JSESSIONID
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "JSESSIONID" {
			JSESSIONID = cookie.Value
		}
	}

	if JSESSIONID == "" {
		return params, errors.New("Can not get JSESSIONID")
	}

	// 正则匹配 HTML 返回的表单字段
	ltReg := regexp.MustCompile("name=\"lt\".+value=\"(.+)\"")
	executionReg := regexp.MustCompile("name=\"execution\".+value=\"(.+)\"")
	_eventIdReg := regexp.MustCompile("name=\"_eventId\".+value=\"(.+)\"")

	bodyStr := string(body)

	ltArr := ltReg.FindStringSubmatch(bodyStr)
	if len(ltArr) != 2 {
		return params, errors.New("Can not get form paramater: lt")
	}
	lt = ltArr[1]

	execArr := executionReg.FindStringSubmatch(bodyStr)
	if len(execArr) != 2 {
		return params, errors.New("Can not get form paramater: execution")
	}
	execution = execArr[1]

	_eventIdArr := _eventIdReg.FindStringSubmatch(bodyStr)
	if len(_eventIdArr) != 2 {
		return params, errors.New("Can not get form paramater: _eventId")
	}
	_eventId = _eventIdArr[1]

	params.lt = lt
	params.execution = execution
	params._eventId = _eventId
	params.submit = "LOGIN"
	params.JSESSIONID = JSESSIONID

	return params, nil
}
