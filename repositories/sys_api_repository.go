package repositories

import (
	"fmt"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/global"
	"gorm.io/gorm"
	"strings"
)

type SysApiRepository struct {
	db *gorm.DB
}

// NewSysApiRepository 构造方法
func NewSysApiRepository() *SysApiRepository {
	return &SysApiRepository{db: global.DB}
}

// Create 创建API信息
func (r *SysApiRepository) Create(api *entities.SysApi, rule *entities.CasbinRule) RepositoryResult {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// API 描述入库
		if err := tx.Create(api).Error; err != nil {
			return err
		}
		// casbin 策略入库
		if err := tx.Create(rule).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return RepositoryResult{}
	}
	return RepositoryResult{Result: api}
}

// FindOneByUniqueKey 根据联合唯一索引查询出一条记录
func (r *SysApiRepository) FindOneByUniqueKey(api *entities.SysApi) RepositoryResult {
	var condition = &entities.SysApi{
		Path:   api.Path,
		Group:  api.Group,
		Method: api.Method,
	}
	var row entities.SysApi
	err := r.db.Where(condition).Take(&row).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: &row}
}

// Pagination 分页构造器
func (r *SysApiRepository) Pagination(pagination *entities.PageReq) RepositoryResult {
	var (
		apis           []entities.SysApi
		totalRows      int64
		err            error
		orderCondition = "created_at desc"
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
	rows := query.Limit(pagination.PageSize).Offset(pagination.Offset()).Find(&apis)
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
		List:       apis,
	}
	return RepositoryResult{Result: pageResult}
}
