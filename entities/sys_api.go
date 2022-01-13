package entities

import (
	"github.com/voyager-go/start-go-api/entities/internal"
)

type SysApi internal.SysApi

// SysApiServiceCreateReq 创建输入参数
type SysApiServiceCreateReq struct {
	Path        string `json:"path" binding:"required"`
	Description string `json:"description" binding:"required"`
	Group       string `json:"group" binding:"required"`
	Method      int8   `json:"method" binding:"required"`
	IsUse       *int8  `json:"is_use"`
	RoleId      uint64 `json:"role_id" binding:"required"`
}
