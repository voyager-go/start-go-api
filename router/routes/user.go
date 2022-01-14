package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/modules/system/api"
)

// InitUserGroup
//@author: [张文杰](https://github.com/voyager-go)
//@slogan    岁岁平，岁岁安，岁岁平安
//@description:
//@create_date: 2022/1/14
//@create_time: 4:59 下午
//@param: r *gin.RouterGroup
//@return: router gin.IRoutes
func InitUserGroup(r *gin.RouterGroup) (router gin.IRoutes) {
	userGroup := r.Group("")
	{
		userGroup.POST("/user/logout", api.User.Logout)
		userGroup.GET("/user/:id", api.User.Show)
		userGroup.PUT("/user/:id", api.User.Update)
		userGroup.POST("/user", api.User.Create)
	}
	return userGroup
}
