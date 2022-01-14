package entities

import "github.com/voyager-go/start-go-api/entities/internal"

type SysRoleMenu internal.SysRoleMenu

// SysRoleMenuServiceCreateReq 创建关联输入参数
type SysRoleMenuServiceCreateReq struct {
	RoleId  uint64   `json:"role_id"`
	MenuIds []uint64 `json:"menu_ids"`
}
