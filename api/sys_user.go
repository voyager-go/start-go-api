package api

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/pkg/response"
)

type SysUserApi struct{}

var SysUser = SysUserApi{}

func (u *SysUserApi) Find(ctx *gin.Context) {
	response.Ok(ctx)
}
