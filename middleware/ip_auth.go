package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/config"
	"net/http"
)

// IpAuth
//@author: [张文杰](https://github.com/voyager-go)
//@slogan    岁岁平，岁岁安，岁岁平安
//@description: 检查ip是否在白名单中
//@create_date: 2022/1/14
//@create_time: 6:07 下午
//@return: gin.HandlerFunc
func IpAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientIp := ctx.ClientIP()
		flag := false
		for _, value := range config.AllowIpList {
			if value == "*" || clientIp == value {
				flag = true
				break
			}
		}
		if !flag {
			ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("%s 不在ip白名单中", clientIp))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
