package controller

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/middleware"
	"github.com/raiki02/EG/tools"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// user这边操作数据库不频繁，不接入redis，如果有神人喜欢一直换名字和头像当我没说。
type UserControllerHdl interface {
	Login(context.Context) gin.HandlerFunc
	Logout(context.Context) gin.HandlerFunc
	NewUser(context.Context, string, string) error
	GetUserInfo(context.Context) gin.HandlerFunc
	CheckToken(context.Context) gin.HandlerFunc
	UpdateAvatar(context.Context) gin.HandlerFunc
	UpdateUsername(context.Context) gin.HandlerFunc
}

type UserController struct {
	e      *gin.Engine
	udh    dao.UserDAOHdl
	cSvc   ccnuService
	jwtHdl middleware.ClaimsHdl
}

func NewUserController(e *gin.Engine, udh dao.UserDAOHdl, cSvc ccnuService, jh middleware.ClaimsHdl) UserControllerHdl {
	return &UserController{
		e:      e,
		udh:    udh,
		cSvc:   cSvc,
		jwtHdl: jh,
	}
}

func (uc *UserController) UpdateAvatar(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO 更新头像
	}
}

func (uc *UserController) UpdateUsername(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO 更新用户名
	}
}

func (uc *UserController) Login(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		//根据一站式账号登录，登陆后返回token
		studentid := c.PostForm("studentid")
		password := c.PostForm("password")

		//首次登录要初始化信息，全部使用默认
		if !uc.udh.CheckUserExist(c.Request.Context(), tools.StrToInt(studentid)) {
			uc.udh.Insert(c, studentid, password)
		}

		success, err := uc.cSvc.Login(context.Background(), studentid, password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if !success {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "登录失败",
			})
			return
		}
		//token操作
		token := uc.jwtHdl.GenToken(c.Request.Context(), studentid)
		err = uc.jwtHdl.StoreInRedis(c.Request.Context(), studentid, token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token存储失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功: " + studentid,
		})

	}
}

func (uc *UserController) Logout(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		err := uc.jwtHdl.ClearToken(c.Request.Context(), token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "登出失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "登出成功",
		})
	}

}

func (uc *UserController) NewUser(ctx context.Context, sid, pwd string) error {
	return uc.udh.Insert(ctx, sid, pwd)
}

func (uc *UserController) GetUserInfo(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO 我的页面，返回信息，头像去图床找，没有就用默认
		studentid := c.Query("studentid")
		user, err := uc.udh.FindUserById(c.Request.Context(), studentid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "用户不存在",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":   "获取成功",
			"studentid": user.StudentId,
			"username":  user.Name,
			"avatar":    user.Avatar,
		})
	}
}

func (uc *UserController) CheckToken(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO 检查token，过期就返回401，前端重新登录
		token := c.GetHeader("Authorization")
		err := uc.jwtHdl.CheckToken(c.Request.Context(), token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token过期",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "token有效",
		})
	}
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
