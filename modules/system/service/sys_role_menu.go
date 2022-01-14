package service

import (
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/repositories"
)

type sysRoleMenuService struct{}

var SysRoleMenu = new(sysRoleMenuService)

func (s *sysRoleMenuService) Create(args *entities.SysRoleMenuServiceCreateReq) error {
	var roleMenus []entities.SysRoleMenu
	for _, roleId := range args.MenuIds {
		roleMenu := entities.SysRoleMenu{
			RoleId: args.RoleId,
			MenuId: roleId,
		}
		roleMenus = append(roleMenus, roleMenu)
	}
	data := repositories.NewSysRoleMenuRepository().Create(&roleMenus)
	return data.Error
}
