package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/modules/system/api"
)

// InitSyRoleGroup
//@author: [张文杰](https://github.com/voyager-go)
//@slogan    岁岁平，岁岁安，岁岁平安
//@description:
//@create_date: 2022/1/14
//@create_time: 5:21 下午
//@param: r *gin.RouterGroup
//@return: router gin.IRoutes
func InitSyRoleGroup(r *gin.RouterGroup) (router gin.IRoutes) {
	sysApi := r.Group("")
	{
		sysApi.POST("/sys_role", api.SysRole.Create)
		sysApi.POST("/sys_role/menu_setting", api.SysRoleMenu.Create)
	}
	return sysApi
}
