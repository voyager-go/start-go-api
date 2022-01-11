package repositories

import (
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/global"
	"gorm.io/gorm"
)

type SysMenuRepository struct {
	db *gorm.DB
}

// NewSysMenuRepository 构造方法
func NewSysMenuRepository() *SysMenuRepository {
	return &SysMenuRepository{db: global.DB}
}

// Create 创建菜单
func (r *SysMenuRepository) Create(menu *entities.SysBaseMenu) RepositoryResult {
	err := r.db.Create(menu).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: menu}
}

// Update 修改菜单信息
func (r *SysMenuRepository) Update(menu *entities.SysBaseMenu) RepositoryResult {
	err := r.db.Save(menu).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: menu}
}

// DeleteOneById 根据主键删除一条记录
func (r *SysMenuRepository) DeleteOneById(id uint64) RepositoryResult {
	err := r.db.Delete(&entities.SysBaseMenu{Model: global.Model{ID: id}}).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: nil}
}

// FindAll 查询所有
func (r *SysMenuRepository) FindAll() RepositoryResult {
	var menus []entities.SysMenu
	q := r.db.Where("is_use = ?", entities.SysMenuIsUseTrue).Order("sort ASC").Session(&gorm.Session{})
	// 先取出一级菜单
	err := q.Where("level = ?", entities.SysMenuLevelFirst).Find(&menus).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	if len(menus) > 0 {
		// 遍历一级菜单
		for index, menu := range menus {
			var children []entities.SysBaseMenu
			err = q.Where("level = ? AND pid = ?", entities.SysMenuLevelSecond, menu.ID).Find(&children).Error
			if err != nil {
				return RepositoryResult{Error: err}
			}
			// 将子菜单写入到菜单列表中
			menus[index].Children = append(menus[index].Children, children...)
		}
	}
	return RepositoryResult{Result: menus}
}
