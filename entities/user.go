package entities

import (
	"github.com/voyager-go/start-go-api/entities/internal"
	"github.com/voyager-go/start-go-api/global"
	"github.com/voyager-go/start-go-api/pkg/util"
)

type User internal.User

const UserStatusNormal = 1    // 启用
const UserStatusForbidden = 0 // 禁用

// UserServiceCreateReq 创建user输入参数
type UserServiceCreateReq struct {
	Nickname string `json:"nickname" binding:"required"`
	Phone    string `json:"phone"    binding:"required"`
	Password string `json:"password" binding:"required"`
	Status   int8   `json:"status"   binding:"required"`
}

// UserServiceUpdateReq 更新user输入参数
type UserServiceUpdateReq struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Status   int8   `json:"status"`
}

// UserServiceTokenReq 请求发放令牌
type UserServiceTokenReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// UserServicePaginationReq 分页输入参数
type UserServicePaginationReq struct {
	util.Pagination
	Searches []global.Search `json:"searches"`
	Rows     interface{}     `json:"rows"`
}
