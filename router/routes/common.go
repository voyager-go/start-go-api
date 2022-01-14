package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/modules/system/api"
)

//InitPublicCommonRouter
//@author: [sspa](https://github.com/voyager-go)
//@slogan    岁岁平，岁岁安，岁岁平安
//@description: 初始化公用基础路由
//@create_date: 2022/1/14
//@create_time: 4:48 下午
//@param: r *gin.RouterGroup
//@return: router gin.IRoutes
func InitPublicCommonRouter(r *gin.RouterGroup) (router gin.IRoutes) {
	commonRoutes := r.Group("")
	{
		// 健康检查
		commonRoutes.GET("/ping", api.Check.Ping)
		// 授权登录
		commonRoutes.POST("/user/login", api.User.Login)
	}
	return commonRoutes
}
