package util

import "github.com/voyager-go/start-go-api/config"

// Pagination 分页器
type Pagination struct {
	Page     uint64 `json:"page" form:"page"`
	PageSize uint64 `json:"pageSize" form:"pageSize"`
	Total    uint64 `json:"total"`
}

// Secure 分页的默认配置与超限默认值
func (p *Pagination) Secure() *Pagination {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = config.Conf.DefaultPageSize
	}
	if p.PageSize > 500 {
		p.PageSize = config.Conf.MaxPageSize
	}
	return p
}

// Offset 分页的偏移量计算
func (p *Pagination) Offset() uint64 {
	p.Secure()
	return (p.Page - 1) * p.PageSize
}

// Limit 分页的限制条数
func (p *Pagination) Limit() uint64 {
	p.Secure()
	return p.PageSize
}

// NewPagination 分页器构造函数
func NewPagination(page, pageSize uint64) *Pagination {
	pagination := &Pagination{
		Page:     page,
		PageSize: pageSize,
	}
	pagination.Secure()
	return pagination
}
