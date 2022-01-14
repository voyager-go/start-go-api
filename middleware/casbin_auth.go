package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/global"
	"net/http"
)

// CasbinAuth
//@author: [张文杰](https://github.com/voyager-go)
//@slogan    岁岁平，岁岁安，岁岁平安
//@description: RABC 角色权限访问控制
//@create_date: 2022/1/14
//@create_time: 6:06 下午
//@param: ctx *gin.Context
//@return: gin.HandlerFunc
func CasbinAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e := global.Enforcer
		// 获取当前请求的 URI
		obj := ctx.Request.URL.RequestURI()
		act := ctx.Request.Method
		// 获取用户的角色
		//user, err := app.GetLoginUser(ctx)
		sub := "root"
		fmt.Println(sub, obj, act)
		// 判断策略中是否存在
		fmt.Println(e.Enforce(sub, obj, act))
		if ok, _ := e.Enforce(sub, obj, act); ok {
			ctx.JSON(http.StatusOK, gin.H{"message": "Authorize pass ..."})
			ctx.Next()
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Authorize forbidden ..."})
			ctx.Abort()
		}
	}
}
