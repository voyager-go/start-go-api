package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/global/app"
	"github.com/voyager-go/start-go-api/pkg/response"
	"net/http"
)

// NeedLogin
//@author: [张文杰](https://github.com/voyager-go)
//@slogan    岁岁平，岁岁安，岁岁平安
//@description: 登录中间件
//@create_date: 2022/1/14
//@create_time: 6:08 下午
//@param: ctx *gin.Context
func NeedLogin(ctx *gin.Context) {
	token := ctx.GetHeader(config.Conf.TokenKey)
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, response.CodeMap[response.UnAuthed])
		ctx.Abort()
		return
	}
	if _, err := app.ParseUserByToken(token); err != nil {
		ctx.JSON(http.StatusUnauthorized, response.CodeMap[response.UnAuthed])
		ctx.Abort()
		return
	}
	ctx.Next()
}
