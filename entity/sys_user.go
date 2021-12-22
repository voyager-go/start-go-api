package entity

import "github.com/voyager-go/start-go-api/entity/internal"

type SysUser internal.SysUser

const SysUserStatusNormal = 1    // 启用
const SysUserStatusForbidden = 0 // 禁用

// SysUserServiceCreateReq 创建sys_user输入参数
type SysUserServiceCreateReq struct {
	Nickname string `json:"nickname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Status   int8   `json:"status" binding:"required"`
}

// SysUserServiceUpdateReq 更新sys_user输入参数
type SysUserServiceUpdateReq struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Status   int8   `json:"status"`
}

// SysUserServiceTokenReq 请求发放令牌
type SysUserServiceTokenReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
