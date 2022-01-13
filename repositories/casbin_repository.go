package repositories

import (
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/global"
	"gorm.io/gorm"
)

type CasbinRuleRepository struct {
	db *gorm.DB
}

func NewCasbinRepository() *CasbinRuleRepository {
	return &CasbinRuleRepository{db: global.DB}
}

func (r *CasbinRuleRepository) Create(rule *entities.CasbinRule) RepositoryResult {
	err := r.db.Create(rule).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}
	return RepositoryResult{Result: rule}
}
