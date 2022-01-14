package api

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/modules/system/service"
	"github.com/voyager-go/start-go-api/pkg/response"
	"github.com/voyager-go/start-go-api/pkg/validator"
)

type sysRoleMenuApi struct{}

var SysRoleMenu = sysRoleMenuApi{}

// Create
// @Summary 为角色关联菜单
// @Schemes
// @Description
// @Tags 角色管理
// @Accept application/json
// @Param Authorization   header string true "Authorization"
// @Param params body entities.SysRoleMenuServiceCreateReq true "角色关联菜单所需信息"
// @Success 200 {object} response.JsonResponse
// @Route
func (s sysRoleMenuApi) Create(ctx *gin.Context) {
	var args entities.SysRoleMenuServiceCreateReq
	err := ctx.ShouldBindJSON(&args)
	if err != nil {
		errs := validator.Translate(err)
		response.FailWithMessage(ctx, errs[0])
		return
	}
	err = service.SysRoleMenu.Create(&args)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
