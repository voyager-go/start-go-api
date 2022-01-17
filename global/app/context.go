package app

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/global"
	"github.com/voyager-go/start-go-api/pkg/auth"
	"strconv"
)

type LoginUser struct {
	entities.User
}

type TokenPayload struct {
	UserId int64 `json:"id"`
}

// ParseUserByToken 通过Token 解析用户
func ParseUserByToken(token string) (TokenPayload, error) {
	user := TokenPayload{}
	if token == "" {
		return user, errors.New("token 为空")
	}
	jwtPayload, err := auth.ParseJwtToken(token, config.Conf.Server.JwtSecret)
	if err != nil {
		return user, err
	}
	byteSlice, err := json.Marshal(jwtPayload.User)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(byteSlice, &user)
	if err != nil {
		return user, err
	}
	if user.UserId == 0 {
		return user, errors.New("非法登录")
	}
	_, err = global.Redis.Get(context.Background(), config.Conf.Redis.LoginPrefix+strconv.FormatInt(user.UserId, 10)).Result()
	if err != nil {
		return TokenPayload{}, errors.New("会话过期，请重新登录")
	}
	return user, nil
}

// GetLoginUser 获取当前登录的用户信息
func GetLoginUser(ctx *gin.Context) (LoginUser, error) {
	info, err := ParseUserByToken(ctx.GetHeader(config.Conf.Server.TokenKey))
	if err != nil {
		return LoginUser{}, err
	}
	Uid := info.UserId
	// 从Redis中查询
	result, err := global.Redis.Get(context.Background(), config.Conf.Redis.LoginPrefix+strconv.FormatInt(Uid, 10)).Result()
	if err != nil {
		return LoginUser{}, err
	}
	var user entities.User
	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		return LoginUser{}, err
	}
	return LoginUser{user}, nil
}
