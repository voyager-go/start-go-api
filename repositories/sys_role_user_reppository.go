package repositories

import (
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/global"
	"gorm.io/gorm"
)

type SysRoleUserRepository struct {
	db *gorm.DB
}

// NewSysRoleUserRepository 构造方法
func NewSysRoleUserRepository() *SysRoleUserRepository {
	return &SysRoleUserRepository{db: global.DB}
}

// FindAllByUserId 查询角色
func (r *SysRoleUserRepository) FindAllByUserId(userId uint64) RepositoryResult {
	var rows []entities.SysRoleUser
	condition := &entities.SysRoleUser{UserId: userId}
	err := r.db.Select("role_id").Where(condition).Find(&rows).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: rows}
}
