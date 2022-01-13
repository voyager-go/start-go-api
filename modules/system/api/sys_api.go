package api

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/modules/system/service"
	"github.com/voyager-go/start-go-api/pkg/response"
	"github.com/voyager-go/start-go-api/pkg/validator"
)

type sysApi struct{}

var SysApi = sysApi{}

// Create
// @Summary 创建一条API记录
// @Schemes
// @Description
// @Tags API操作
// @Accept application/json
// @Param Authorization header string true "Authorization"
// @Param userInfo body entities.SysApiServiceCreateReq true "API基础信息"
// @Success 200 {string} response.Ok
// @Router /sys_api [post]
func (u sysApi) Create(ctx *gin.Context) {
	var args entities.SysApiServiceCreateReq
	err := ctx.ShouldBindJSON(&args)
	if err != nil {
		errs := validator.Translate(err) // 验证器返回错误信息后，翻译成中文
		response.FailWithMessage(ctx, errs[0])
		return
	}
	err = service.SysApi.Create(&args)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

// List
// @Summary 获取api分页数据
// @Schemes
// @Description
// @Tags API操作
// @Accept application/json
// @Param Authorization header string true "Authorization"
// @Param page query int true "当前页"
// @Param pageSize query int true "每页条目"
// @Param searches query []global.Search false "筛选条件"
// @Success 200 {string} response.Ok
// @Router /sys_api/list [get]
func (u sysApi) List(ctx *gin.Context) {
	var args entities.PageReq
	err := ctx.ShouldBindQuery(&args)
	if err != nil {
		errs := validator.Translate(err) // 验证器返回错误信息后，翻译成中文
		response.FailWithMessage(ctx, errs[0])
		return
	}
	result, err := service.SysApi.Page(&args)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
