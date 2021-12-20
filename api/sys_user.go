package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/voyager-go/start-go-api/dao"
	"github.com/voyager-go/start-go-api/entity"
	"github.com/voyager-go/start-go-api/global/app"
	"github.com/voyager-go/start-go-api/pkg/response"
	"github.com/voyager-go/start-go-api/service"
)

type SysUserApi struct{}

var SysUser = new(SysUserApi)

// Find
// @Summary 根据用户编号查找一条用户记录
// @Schemes
// @Description
// @Tags 用户操作
// @Param Authorization header string true "Authorization"
// @Param id path int true "用户编号"
// @Success 200 {string} response.OkWithData
// @Router /user/{id} [get]
func (u SysUserApi) Find(ctx *gin.Context) {
	id := gconv.Int64(ctx.Param("id"))
	if id == 0 {
		response.FailWithDetail(ctx, response.RequestParamErr)
		return
	}
	user, err := dao.SysUser.FindOneById(id, true)
	if err != nil {
		response.FailWithMessage(ctx, "未查询到对应用户或该用户状态异常")
		return
	}
	response.OkWithData(ctx, user)
}

// Create
// @Summary 创建一条用户记录
// @Schemes
// @Description
// @Tags 用户操作
// @Accept application/json
// @Param Authorization header string true "Authorization"
// @Param userInfo body entity.SysUserServiceCreateReq true "用户基础信息"
// @Success 200 {string} response.Ok
// @Router /user [post]
func (u SysUserApi) Create(ctx *gin.Context) {
	var args entity.SysUserServiceCreateReq
	err := ctx.ShouldBindJSON(&args)
	if err != nil {
		response.FailWithDetail(ctx, response.RequestParamErr)
		return
	}
	err = service.User.CreateUser(args)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

// Update
// @Summary 更新用户信息
// @Schemes
// @Description
// @Tags 用户
// @Param Authorization header string true "Authorization"
// @Param userInfo body entity.SysUserServiceUpdateReq true "需要更新的用户信息"
// @Success 200 {string} response.Ok
// @Router /user [put]
func (u SysUserApi) Update(ctx *gin.Context) {
	var args entity.SysUserServiceUpdateReq
	err := ctx.ShouldBindJSON(&args)
	if err != nil {
		response.FailWithDetail(ctx, response.RequestParamErr)
		return
	}
	err = service.User.UpdateUser(args)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

// Login
// @Summary 用户授权登录
// @Schemes
// @Description
// @Accept application/json
// @Tags 用户操作
// @Param userInfo body entity.SysUserServiceTokenReq true "手机号和密码"
// @Success 200 {string} response.OkWithData
// @Router /user/auth [post]
func (u SysUserApi) Login(ctx *gin.Context) {
	var args entity.SysUserServiceTokenReq
	err := ctx.ShouldBindJSON(&args)
	if err != nil {
		response.FailWithDetail(ctx, response.RequestParamErr)
		return
	}
	token, err := service.User.Login(args)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, token)
}

// Logout
// @Summary 退出会话
// @Schemes
// @Description
// @Tags 用户操作
// @Accept application/json
// @Param Authorization   header string true "Authorization"
// @Success 200 {string} response.Ok
// @Router /user [delete]
func (u SysUserApi) Logout(ctx *gin.Context) {
	user, err := app.GetLoginUser(ctx)
	if err != nil {
		response.FailWithDetail(ctx, response.AuthExpired)
		return
	}
	err = service.User.Logout(user.ID)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
