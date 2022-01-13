package repositories

import (
	"github.com/voyager-go/start-go-api/global"
	"gorm.io/gorm"
)

type SysRoleMenuRepository struct {
	db *gorm.DB
}

func NewSysRoleMenuRepository() *SysRoleMenuRepository {
	return &SysRoleMenuRepository{db: global.DB}
}

func (r SysRoleMenuRepository) Create() {

}
