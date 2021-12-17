package entity

import "github.com/voyager-go/start-go-api/entity/internal"

type SysUser internal.SysUser

const SysUserStatusNormal = 1    // 启用
const SysUserStatusForbidden = 0 // 禁用

// SysUserServiceCreateReq 创新sys_user输入参数
type SysUserServiceCreateReq struct {
	Nickname string
	Phone    string
	Password string
	Status   int8
}

// SysUserServiceChangeStatusReq 更新用户状态
type SysUserServiceChangeStatusReq struct {
	Id     int64 `json:"id"`
	Status int8  `json:"status"`
}
