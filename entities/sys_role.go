package entities

import "github.com/voyager-go/start-go-api/entities/internal"

type SysRole internal.SysRole

type SysRoleServiceCreateReq struct {
	Name  string  `json:"name" binding:"required"`
	Pid   *uint64 `json:"pid"`
	IsUse *int8   `json:"is_use"`
}
