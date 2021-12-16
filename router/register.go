package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/api"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/middleware"
	"net/http"
)

func Register() *gin.Engine {
	gin.SetMode(config.Conf.Mode)
	router := gin.New()
	// 404 处理
	router.NoRoute(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		fmt.Println(path)
		method := ctx.Request.Method
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("%s %s not found", method, path))
	})
	// 健康检查
	router.GET("/ping", api.Check.Ping)
	// 路由分组
	var (
		publicMiddleware = []gin.HandlerFunc{
			cors.Default(),
			middleware.IpAuth(),
		}
		// 用户组
		user = router.Group("/user", publicMiddleware...)
	)
	user.GET("/find", api.SysUser.Find)

	return router
}
