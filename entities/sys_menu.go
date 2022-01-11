package entities

import "github.com/voyager-go/start-go-api/entities/internal"

type SysBaseMenu internal.SysBaseMenu

type SysMenu struct {
	SysBaseMenu
	Children []SysBaseMenu `gorm:"-" json:"children"`
}

func (SysBaseMenu) TableName() string {
	return "sys_menu"
}

const (
	SysMenuLevelFirst = 1 // 一级菜单

	SysMenuLevelSecond = 2 // 二级菜单

	SysMenuIsUseFalse = 0 // 禁用

	SysMenuIsUseTrue = 1 // 启用
)

// SysMenuServiceReq 创建菜单输入参数
type SysMenuServiceReq struct {
	Name      string `json:"name" binding:"required"`
	Pid       *int64 `json:"pid" binding:"required"`
	IsUse     *int8  `json:"is_use" binding:"required"`
	Level     int8   `json:"level" binding:"required"`
	Sort      int8   `json:"sort" binding:"gt=0"`
	Icon      string `json:"icon"`
	UniqueKey string `json:"unique_key" binding:"required"`
}
