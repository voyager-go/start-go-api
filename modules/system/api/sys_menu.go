package api

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/modules/system/service"
	"github.com/voyager-go/start-go-api/pkg/response"
	"github.com/voyager-go/start-go-api/pkg/validator"
)

type sysMenuApi struct{}

var Menu sysMenuApi

// Create
// @Summary 创建一条用户记录
// @Schemes
// @Description
// @Tags 菜单操作
// @Accept application/json
// @Param Authorization   header string true "Authorization"
// @Param params body entities.SysMenuServiceReq true "菜单基础信息"
// @Success 200 {string} response.Ok
// @Router /sys_menu [post]
func (m sysMenuApi) Create(ctx *gin.Context) {
	var args entities.SysMenuServiceReq
	err := ctx.ShouldBindJSON(&args)
	if err != nil {
		errs := validator.Translate(err)
		response.FailWithMessage(ctx, errs[0])
		return
	}
	err = service.SysMenu.Create(&args)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

// List
// @Summary 查询菜单列表
// @Schemes
// @Description
// @Tags 菜单操作
// @Accept application/json
// @Param Authorization   header string true "Authorization"
// @Success 200 {string} response.OkWithData
// @Router /sys_menu [get]
func (m sysMenuApi) List(ctx *gin.Context) {
	list, err := service.SysMenu.FindList()
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, list)
}
