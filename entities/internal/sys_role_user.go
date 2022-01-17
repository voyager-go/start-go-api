package internal

import "github.com/voyager-go/start-go-api/global"

// SysRoleUser 角色与用户关联
type SysRoleUser struct {
	global.Model
	RoleId uint64 `gorm:"column:role_id;type:bigint(20) unsigned;NOT NULL;comment:角色编号" json:"role_id"` // 角色编号
	UserId uint64 `gorm:"column:user_id;type:bigint(20) unsigned;NOT NULL;comment:用户编号" json:"user_id"` // 用户编号
}
