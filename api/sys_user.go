package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/voyager-go/start-go-api/dao"
	"github.com/voyager-go/start-go-api/entity"
	"github.com/voyager-go/start-go-api/pkg/response"
	"github.com/voyager-go/start-go-api/service"
)

type SysUserApi struct{}

var SysUser = SysUserApi{}

func (u *SysUserApi) Find(ctx *gin.Context) {
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

func (u SysUserApi) ChangeStatus(ctx *gin.Context) {
	var data entity.SysUserServiceChangeStatusReq
	err := ctx.ShouldBindJSON(&data)
	if data.Id == 0 {
		response.FailWithDetail(ctx, response.RequestParamErr)
		return
	}
	if err != nil {
		response.FailWithDetail(ctx, response.RequestParamErr)
		return
	}
	err = service.User.ChangeUserStatus(data)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (u *SysUserApi) Login(ctx *gin.Context) {
	response.OkWithData(ctx, gin.H{"token": "saklsajlsdajldkaslj"})
}
