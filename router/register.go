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
	"github.com/voyager-go/start-go-api/router/routes"
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
	// 路由分组
	var (
		publicMiddleware = []gin.HandlerFunc{
			cors.Default(),
			middleware.IpAuth(),
		}
		apiGroup     = router.Group("/api", publicMiddleware...)
		apiNeedLogin = router.Group("/api", append(publicMiddleware, middleware.NeedLogin, middleware.CasbinAuth())...)
	)
	// 公用组
	routes.InitPublicCommonRouter(apiGroup)
	// 角色组
	routes.InitSyRoleGroup(apiNeedLogin)
	// 用户组
	routes.InitUserGroup(apiNeedLogin)
	// 菜单组
	routes.InitSysMenuGroup(apiNeedLogin)
	// API组
	routes.InitSysApiGroup(apiNeedLogin)
	return router
}
