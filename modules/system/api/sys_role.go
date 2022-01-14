package api

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/modules/system/service"
	"github.com/voyager-go/start-go-api/pkg/response"
	"github.com/voyager-go/start-go-api/pkg/validator"
)

type sysRoleApi struct{}

var SysRole = sysRoleApi{}

// Create
// @Summary 新增一个角色
// @Schemes
// @Description
// @Tags 角色管理
// @Accept application/json
// @Param Authorization   header string true "Authorization"
// @Param role body entities.SysRoleServiceCreateReq true "角色基础信息"
// @Success 200 {object} response.JsonResponse
// @Router /sys_role [post]
func (s sysRoleApi) Create(ctx *gin.Context) {
	var role entities.SysRoleServiceCreateReq
	err := ctx.ShouldBindJSON(&role)
	if err != nil {
		errs := validator.Translate(err) // 验证器返回错误信息后，翻译成中文
		response.FailWithMessage(ctx, errs[0])
		return
	}
	err = service.SysRole.Create(&role)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
