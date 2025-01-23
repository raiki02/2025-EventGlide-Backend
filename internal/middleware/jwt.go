package middleware

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

// userhandler 要操作这个 管理token
type ClaimsHdl interface {
	GenToken(context.Context, string) string
	StoreInRedis(context.Context, string, string) error
	CheckToken(context.Context, string) error
	ClearToken(context.Context, string) error
}
type Claims struct {
	rdb    *redis.Client
	jwtKey []byte
}

func NewClaimsHdl(rdb *redis.Client) ClaimsHdl {
	jwtKey := viper.GetString("jwt.key")
	return Claims{
		jwtKey: []byte(jwtKey),
		rdb:    rdb,
	}
}

func (c Claims) GenToken(ctx context.Context, sid string) string {
	claims := jwt.RegisteredClaims{
		ID:        uuid.New().String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		Subject:   sid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString(c.jwtKey)
	return t
}

func (c Claims) StoreInRedis(ctx context.Context, sid string, token string) error {
	//把token解析出对应id 存入redis中
	id := c.parseTokenId(token)
	key := "token:" + sid
	err := c.rdb.Set(ctx, key, id, time.Hour*72).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c Claims) CheckToken(ctx context.Context, token string) error {
	id := c.parseTokenId(token)
	if id == "" {
		return errors.New("token is invalid")
	}
	_, err := c.rdb.Get(ctx, id).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c Claims) ClearToken(ctx context.Context, sid string) error {
	err := c.rdb.Del(ctx, sid).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c Claims) parseTokenId(token string) string {
	t, _ := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return c.jwtKey, nil
	})
	if c, ok := t.Claims.(*jwt.RegisteredClaims); ok && t.Valid {
		return c.ID
	}
	return ""
}
