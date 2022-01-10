package service

import (
	"github.com/gogf/gf/util/gconv"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/repositories"
)

type sysMenu struct{}

var SysMenuService sysMenu

// Create 创建一条记录
func (m *sysMenu) Create(req *entities.SysMenuServiceReq) error {
	var menu entities.SysBaseMenu
	err := gconv.Struct(req, &menu)
	if err != nil {
		return err
	}
	return repositories.NewSysMenuRepository().Create(&menu).Error
}

// FindList 查询全部记录
func (m *sysMenu) FindList() ([]entities.SysMenu, error) {
	var menus []entities.SysMenu
	result := repositories.NewSysMenuRepository().FindAll()
	if result.Error != nil {
		return menus, result.Error
	}
	return result.Result.([]entities.SysMenu), nil
}