package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/global/app"
	"github.com/voyager-go/start-go-api/pkg/response"
	"net/http"
)

// NeedLogin 登录中间件
func NeedLogin(ctx *gin.Context) {
	token := ctx.GetHeader(config.Conf.TokenKey)
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, response.CodeMap[response.UnAuthed])
		ctx.Abort()
		return
	}
	if _, err := app.ParseUserByToken(token); err != nil {
		ctx.JSON(http.StatusUnauthorized, response.CodeMap[response.UnAuthed]+"，详情为:"+err.Error())
		ctx.Abort()
		return
	}
	ctx.Next()
}
