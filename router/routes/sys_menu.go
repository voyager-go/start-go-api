package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/modules/system/api"
)

// InitSysMenuGroup
//@author: [张文杰](https://github.com/voyager-go)
//@slogan    岁岁平，岁岁安，岁岁平安
//@description:
//@create_date: 2022/1/14
//@create_time: 5:21 下午
//@param: r *gin.RouterGroup
//@return: router gin.IRoutes
func InitSysMenuGroup(r *gin.RouterGroup) (router gin.IRoutes) {
	menuGroup := r.Group("")
	{
		menuGroup.POST("/sys_menu", api.Menu.Create)
		menuGroup.GET("/sys_menu", api.Menu.List)
	}
	return menuGroup
}
