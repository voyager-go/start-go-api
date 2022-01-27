package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/entities/enum"
	"github.com/voyager-go/start-go-api/repositories"
	"strconv"
)

type sysApiService struct{}

var SysApi = new(sysApiService)

// Create 创建一条API信息
func (a *sysApiService) Create(req *entities.SysApiServiceCreateReq) error {
	var api entities.SysApi
	err := copier.Copy(&api, req)
	if err != nil {
		return err
	}
	var repository = repositories.NewSysApiRepository()
	checkResult := repository.FindOneByUniqueKey(&api)
	if row, ok := checkResult.Result.(*entities.SysApi); ok && row.ID > 0 {
		return errors.New("api访问路径、所在分组、及访问方法是联合唯一的")
	}
	ruleArgs := entities.CasbinRuleServiceCreateReq{
		RoleID: strconv.FormatUint(req.RoleId, 10),
		Path:   req.Path,
		Method: enum.MethodType(req.Method).String(),
	}
	var rule entities.CasbinRule
	err = copier.Copy(&rule, ruleArgs)
	if err != nil {
		return err
	}
	return repository.Create(&api, &rule).Error
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
