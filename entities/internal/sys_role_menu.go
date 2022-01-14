package internal

// SysRoleMenu 菜单角色关联表
type SysRoleMenu struct {
	ID     uint64 `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"` // 主键
	RoleId uint64 `gorm:"column:role_id;type:bigint(20) unsigned;NOT NULL" json:"role_id"`         // 角色ID
	MenuId uint64 `gorm:"column:menu_id;type:bigint(20) unsigned;NOT NULL" json:"menu_id"`         // 菜单ID
}
