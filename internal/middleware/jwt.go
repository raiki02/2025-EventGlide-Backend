package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/raiki02/EG/tools"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

type JwtHdl interface {
	GenToken(*gin.Context, string) string
	StoreInRedis(*gin.Context, string, string) error
	CheckToken(*gin.Context, string) error
	ClearToken(*gin.Context, string) error
}
type Jwt struct {
	rdb    *redis.Client
	jwtKey []byte
}

func NewJwt(rdb *redis.Client) *Jwt {
	jwtKey := viper.GetString("jwt.key")
	return &Jwt{
		jwtKey: []byte(jwtKey),
		rdb:    rdb,
	}
}

func (c *Jwt) GenToken(ctx *gin.Context, sid string) string {
	claims := jwt.RegisteredClaims{
		ID:        uuid.New().String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(setTTL())),
		Subject:   sid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(c.jwtKey)
	if err != nil {
		return ""
	}

	return t
}

func (c *Jwt) StoreInRedis(ctx *gin.Context, sid string, token string) error {
	id := c.parseTokenId(token)
	key := "token:" + id
	err := c.rdb.Set(ctx, key, sid, setTTL()).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Jwt) CheckToken(ctx *gin.Context, token string) error {
	id := c.parseTokenId(token)
	if id == "" {
		return errors.New("token is invalid")
	}
	id = "token:" + id
	_, err := c.rdb.Get(ctx, id).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c *Jwt) ClearToken(ctx *gin.Context, token string) error {
	id := c.parseTokenId(token)
	id = "token:" + id
	err := c.rdb.Del(ctx, id).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Jwt) parseTokenId(token string) string {
	if token == "" {
		return ""
	}
	if strings.HasPrefix(token, "Bearer") {
		_, token, _ = strings.Cut(token, " ")
	}
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return c.jwtKey, nil
	})
	if err != nil {
		return ""
	}
	if c, ok := t.Claims.(*jwt.RegisteredClaims); ok && t.Valid {
		return c.ID
	}
	return ""
}

func setTTL() time.Duration {
	ttl, _ := strconv.Atoi(viper.GetString("jwt.ttl"))
	return time.Second * time.Duration(ttl)
}

func (c *Jwt) WrapCheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(401, tools.ReturnMSG(ctx, "token is empty", nil))
			ctx.Abort()
			return
		}
		err := c.CheckToken(ctx, token)
		if err != nil {
			ctx.JSON(401, tools.ReturnMSG(ctx, "token is invalid", nil))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func (c *Jwt) storeSid(token string) error {

}
