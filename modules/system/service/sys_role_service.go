package service

import (
	"github.com/jinzhu/copier"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/repositories"
)

type sysRoleService struct{}

var SysRole = new(sysRoleService)

func (s *sysRoleService) Create(req *entities.SysRoleServiceCreateReq) error {
	var sysRole entities.SysRole
	err := copier.Copy(&sysRole, req)
	if err != nil {
		return err
	}
	return repositories.NewSysRoleRepository().Create(&sysRole).Error
}
