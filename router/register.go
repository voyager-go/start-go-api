package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/docs"
	"github.com/voyager-go/start-go-api/middleware"
	"github.com/voyager-go/start-go-api/modules/system/api"
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
	// swagger 配置
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 健康检查
	router.GET("/ping", api.Check.Ping)
	// 路由分组
	var (
		publicMiddleware = []gin.HandlerFunc{
			cors.Default(),
			middleware.IpAuth(),
		}
		// 用户组
		apiGroup     = router.Group("/api", publicMiddleware...)
		apiNeedLogin = router.Group("/api", append(publicMiddleware, middleware.NeedLogin)...)
	)
	apiGroup.POST("/user/login", api.User.Login)
	apiNeedLogin.POST("/user/logout", api.User.Logout)
	apiNeedLogin.GET("/user/:id", api.User.Show)
	apiNeedLogin.PUT("/user/:id", api.User.Update)
	apiNeedLogin.POST("/user", api.User.Create)

	// 菜单组
	apiNeedLogin.POST("/sys_menu", api.Menu.Create)
	apiNeedLogin.GET("/sys_menu", api.Menu.List)

	// API组
	apiNeedLogin.POST("/sys_api", api.SysApi.Create)
	apiNeedLogin.GET("/sys_api/list", api.SysApi.List)
	return router
}
