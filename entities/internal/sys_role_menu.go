package internal

// SysRoleMenu 菜单角色关联表
type SysRoleMenu struct {
	ID     uint64 `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT;comment:主键" json:"id"` // 主键
	RoleId uint64 `gorm:"column:role_id;type:bigint(20) unsigned;NOT NULL;comment:角色ID" json:"role_id"`       // 角色ID
	MenuId uint64 `gorm:"column:menu_id;type:bigint(20) unsigned;NOT NULL;comment:菜单ID" json:"menu_id"`       // 菜单ID
}
