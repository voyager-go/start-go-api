package service

import (
	"errors"
	"github.com/gogf/gf/util/gconv"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/repositories"
)

type sysApiService struct{}

var SysApi = new(sysApiService)

// Create 创建一条API信息
func (a *sysApiService) Create(req *entities.SysApiServiceCreateReq) error {
	var api entities.SysApi
	err := gconv.Struct(req, &api)
	if err != nil {
		return err
	}
	var repository = repositories.NewSysApiRepository()
	checkResult := repository.FindOneByUniqueKey(&api)
	if row, ok := checkResult.Result.(*entities.SysApi); ok && row.ID > 0 {
		return errors.New("api访问路径、所在分组、及访问方法是联合唯一的")
	}
	return repository.Create(&api).Error
}

// Page 分页信息
func (a *sysApiService) Page(req *entities.PageReq) (entities.PageResult, error) {
	var pageResult entities.PageResult
	result := repositories.NewSysApiRepository().Pagination(req)
	if result.Error != nil {
		return pageResult, result.Error
	}
	pageResult, ok := result.Result.(entities.PageResult)
	if !ok {
		return pageResult, errors.New("分页结果异常")
	}
	return pageResult, nil
}
