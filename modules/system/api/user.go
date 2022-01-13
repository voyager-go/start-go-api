package api

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/global/app"
	"github.com/voyager-go/start-go-api/modules/system/service"
	"github.com/voyager-go/start-go-api/pkg/response"
	"github.com/voyager-go/start-go-api/pkg/validator"
	"strconv"
)

type userApi struct{}

var User = userApi{}

// Login
// @Summary 用户授权登录
// @Schemes
// @Description
// @Accept application/json
// @Tags 用户操作
// @Param userInfo body entities.UserServiceTokenReq true "手机号和密码"
// @Success 200 {string} response.OkWithData
// @Router /user/login [post]
func (u userApi) Login(ctx *gin.Context) {
	var args entities.UserServiceTokenReq
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

// Show
// @Summary 查看用户记录
// @Schemes
// @Description
// @Tags 用户操作
// @Accept application/json
// @Param Authorization header string true "Authorization"
// @Param id path int true "用户编号"
// @Success 200 {string} response.Ok
// @Router /user/{id} [get]
func (u userApi) Show(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.FailWithDetail(ctx, response.RequestParamErr)
		return
	}
	user, err := service.User.FindOne(uint64(userId))
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
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
// @Param userInfo body entities.UserServiceCreateReq true "用户基础信息"
// @Success 200 {string} response.Ok
// @Router /user [post]
func (u userApi) Create(ctx *gin.Context) {
	var args entities.UserServiceCreateReq
	err := ctx.ShouldBindJSON(&args)
	if err != nil {
		errs := validator.Translate(err) // 验证器返回错误信息后，翻译成中文
		response.FailWithMessage(ctx, errs[0])
		return
	}
	err = service.User.Create(&args)
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
// @Tags 用户操作
// @Param Authorization header string true "Authorization"
// @Param id path int true "用户编号"
// @Param userInfo body entities.UserServiceUpdateReq true "需要更新的用户信息"
// @Success 200 {string} response.Ok
// @Router /user/{id} [put]
func (u userApi) Update(ctx *gin.Context) {
	var args entities.UserServiceUpdateReq
	err := ctx.ShouldBindJSON(&args)
	if err != nil {
		response.FailWithDetail(ctx, response.RequestParamErr)
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id == 0 || err != nil {
		response.FailWithDetail(ctx, response.RequestParamErr)
		return
	}
	args.ID = uint64(id)
	err = service.User.Update(args)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

// Logout
// @Summary 退出会话
// @Schemes
// @Description
// @Tags 用户操作
// @Accept application/json
// @Param Authorization   header string true "Authorization"
// @Success 200 {string} response.Ok
// @Router /user/logout [post]
func (u userApi) Logout(ctx *gin.Context) {
	user, err := app.GetLoginUser(ctx)
	if err != nil {
		response.FailWithDetail(ctx, response.AuthExpired)
		return
	}
	err = service.User.Logout(int64(user.ID))
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
