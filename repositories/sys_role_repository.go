package repositories

import (
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/global"
	"gorm.io/gorm"
)

type SysRoleRepository struct {
	db *gorm.DB
}

// NewSysRoleRepository 构造方法
func NewSysRoleRepository() *SysRoleRepository {
	return &SysRoleRepository{db: global.DB}
}

// Create 创建角色
func (r *SysRoleRepository) Create(role *entities.SysRole) RepositoryResult {
	err := r.db.Create(role).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: role}
}

// Update 修改角色信息
func (r *SysRoleRepository) Update(role *entities.SysRole) RepositoryResult {
	err := r.db.Save(role).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: role}
}

// DeleteOneById 根据主键删除一条记录
func (r *SysRoleRepository) DeleteOneById(id uint64) RepositoryResult {
	err := r.db.Delete(&entities.SysRole{Model: global.Model{ID: id}}).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: nil}
}

// FindAll 查询所有
func (r *SysRoleRepository) FindAll() RepositoryResult {
	var roles []entities.SysRole
	err := r.db.Find(&roles).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: roles}
}
