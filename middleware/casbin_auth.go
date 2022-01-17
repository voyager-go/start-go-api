package middleware

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/global"
	"github.com/voyager-go/start-go-api/global/app"
	"net/http"
	"strconv"
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
		adapter, _ := gormadapter.NewAdapterByDB(global.DB)
		e, _ := casbin.NewEnforcer("rbac_model.conf", adapter)
		err := e.LoadPolicy()
		// 获取当前请求的 URI
		obj := ctx.Request.URL.RequestURI()
		act := ctx.Request.Method
		// 获取用户的角色
		user, err := app.GetLoginUser(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "权限异常")
			ctx.Abort()
		}
		var flag = false
		for _, sub := range user.RoleIds {
			// 判断策略中是否存在
			subStr := strconv.FormatUint(sub, 10)
			if ok, _ := e.Enforce(subStr, obj, act); ok {
				flag = true
			}
		}
		if flag {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusOK, "该用户无此权限")
			ctx.Abort()
		}
	}
}
