package entities

import (
	"github.com/voyager-go/start-go-api/entities/internal"
)

type BaseUser internal.BaseUser

type User struct {
	BaseUser
	RoleIds []uint64 `gorm:"-" json:"role_ids"`
}

// UserServiceCreateReq 创建user输入参数
type UserServiceCreateReq struct {
	Nickname string `json:"nickname" binding:"required"`
	Phone    string `json:"phone"    binding:"required"`
	Password string `json:"password" binding:"required"`
	Status   *int8  `json:"status"   binding:"required"`
}

// UserServiceUpdateReq 更新user输入参数
type UserServiceUpdateReq struct {
	ID       uint64 `json:"id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Status   *int8  `json:"status"`
}

// UserServiceTokenReq 请求发放令牌
type UserServiceTokenReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
