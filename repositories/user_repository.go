package repositories

import (
	"errors"
	"fmt"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/global"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 构造查询对象
func NewUserRepository() *UserRepository {
	return &UserRepository{db: global.DB}
}

// Create 根据entities.User模型创建一条记录
func (r *UserRepository) Create(user *entities.User) RepositoryResult {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return RepositoryResult{Error: err}
	}
	user.Password = string(fromPassword)
	err = r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(user).Error
		if err != nil {
			return err
		}
		userRoles := make([]entities.SysRoleUser, len(user.RoleIds))
		for _, roleId := range user.RoleIds {
			record := entities.SysRoleUser{
				RoleId: roleId,
				UserId: user.ID,
			}
			userRoles = append(userRoles, record)
		}
		err = tx.Create(&userRoles).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: user}
}

// Update 根据entities.User模型更新一条记录
func (r *UserRepository) Update(user *entities.User) RepositoryResult {
	err := r.db.Save(user).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: user}
}

// FindOneById 通过主键ID查询一条记录
func (r *UserRepository) FindOneById(id uint64) RepositoryResult {
	var user entities.User
	condition := &entities.User{
		BaseUser: entities.BaseUser{
			Model: global.Model{
				ID: id,
			},
		},
	}
	err := r.db.Where(condition).Take(&user).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: &user}
}

// FindAll 查询所有
func (r *UserRepository) FindAll() RepositoryResult {
	var users []entities.User
	err := r.db.Find(&users).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: users}
}

// DeleteOneById 根据主键删除一条记录
func (r *UserRepository) DeleteOneById(id uint64) RepositoryResult {
	condition := &entities.User{
		BaseUser: entities.BaseUser{
			Model: global.Model{
				ID: id,
			},
		},
	}
	err := r.db.Delete(condition).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: nil}
}

// FindOneByPhone 通过手机号查询一条记录
func (r *UserRepository) FindOneByPhone(phone string) RepositoryResult {
	if phone == "" {
		return RepositoryResult{Error: errors.New("手机号不得为空")}
	}
	var user entities.User
	err := r.db.Where(&entities.User{
		BaseUser: entities.BaseUser{
			Phone: phone,
		},
	}).Take(&user).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: &user}
}

// CheckPassword 判断密码是否匹配
func (r *UserRepository) CheckPassword(hashPassword, password string) RepositoryResult {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return RepositoryResult{Error: err}
}

// Pagination 分页构造器
func (r *UserRepository) Pagination(pagination *entities.PageReq) RepositoryResult {
	var (
		users          []entities.User
		totalRows      int64
		err            error
		orderCondition = "created_at desc, nickname"
	)
	query := r.db.Order(orderCondition)
	// 生成WHERE条件
	if pagination.Searches != nil {
		for _, value := range pagination.Searches {
			column := value.Column
			action := value.Action
			needle := value.Needle
			switch action {
			case "equals":
				whereCondition := fmt.Sprintf("%s = ?", column)
				query = query.Where(whereCondition, needle)
				break
			case "contains":
				whereCondition := fmt.Sprintf("%s LIKE ?", column)
				query = query.Where(whereCondition, "%"+needle+"%")
				break
			case "in":
				whereCondition := fmt.Sprintf("%s IN (?)", column)
				querySlice := strings.Split(needle, ",")
				query = query.Where(whereCondition, querySlice)
				break
			}
		}
	}
	rows := query.Limit(pagination.PageSize).Offset(pagination.Offset()).Find(&users)
	err = rows.Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	// 总条目
	err = query.Count(&totalRows).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	// 总条目数量
	pagination.Total = int(totalRows)
	// 返回结构
	pageResult := entities.PageResult{
		Pagination: pagination.Pagination,
		List:       users,
	}
	return RepositoryResult{Result: pageResult}
}
