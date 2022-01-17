package internal

import "github.com/voyager-go/start-go-api/global"

// SysBaseMenu 菜单表
type SysBaseMenu struct {
	global.Model
	Name      string  `gorm:"column:name;type:varchar(30);NOT NULL;comment:菜单名称" json:"name"`                            // 菜单名称
	Pid       *uint64 `gorm:"column:pid;type:bigint(20) unsigned;default:0;NOT NULL;comment:父级ID" json:"pid"`            // 父级ID
	IsUse     *int8   `gorm:"column:is_use;type:tinyint(1);default:1;NOT NULL;comment:是否可用 默认1表示启用 0表示禁用" json:"is_use"` // 是否可用 默认1表示启用 0表示禁用
	Level     int8    `gorm:"column:level;type:tinyint(1);default:1;NOT NULL;comment:菜单级别 默认是1级菜单" json:"level"`         // 菜单级别 默认是1级菜单
	Sort      int8    `gorm:"column:sort;type:tinyint(1);NOT NULL;comment:排序编号" json:"sort"`                             // 排序编号
	Icon      string  `gorm:"column:icon;type:varchar(100);comment:图标ICON" json:"icon"`                                  // 图标ICON
	UniqueKey string  `gorm:"column:unique_key;type:varchar(30);NOT NULL;comment:唯一描述" json:"unique_key"`                // 唯一描述
}
