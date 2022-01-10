package internal

import "github.com/voyager-go/start-go-api/global"

// SysRole 角色表
type SysRole struct {
	global.Model
	Role  string `gorm:"column:role;type:varchar(60);NOT NULL" json:"role"`              // 角色名称
	Pid   int64  `gorm:"column:pid;type:bigint(20) unsigned;NOT NULL" json:"pid"`        // 父级ID
	IsUse *int8  `gorm:"column:is_use;type:tinyint(1);default:1;NOT NULL" json:"is_use"` // 是否可用 默认1表示启用 0表示禁用
}