package service

import (
	"github.com/gogf/gf/util/gconv"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/repositories"
)

type sysRoleService struct{}

var SysRole = new(sysRoleService)

func (s *sysRoleService) Create(req *entities.SysRoleServiceCreateReq) error {
	var sysRole entities.SysRole
	err := gconv.Struct(req, &sysRole)
	if err != nil {
		return err
	}
	return repositories.NewSysRoleRepository().Create(&sysRole).Error
}
