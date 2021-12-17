package service

import (
	"errors"
	"github.com/voyager-go/start-go-api/dao"
	"github.com/voyager-go/start-go-api/entity"
	"github.com/voyager-go/start-go-api/global"
	"github.com/voyager-go/start-go-api/pkg/util"
)

type SysUserService struct{}

var User = SysUserService{}

// CreateUser 创建用户
func (u *SysUserService) CreateUser(r *entity.SysUserServiceCreateReq) error {
	if u.CheckPhoneExists(r.Phone) {
		return errors.New("该手机号已经存在")
	}
	statusSlice := []int{entity.SysUserStatusForbidden, entity.SysUserStatusNormal}
	if !util.InIntSlice(int(r.Status), statusSlice) {
		return errors.New("用户状态有误")
	}
	return global.DB.Save(&r).Error
}

// ChangeUserStatus 修改用户状态
func (u *SysUserService) ChangeUserStatus(data entity.SysUserServiceChangeStatusReq) error {
	statusSlice := []int{entity.SysUserStatusForbidden, entity.SysUserStatusNormal}
	if !util.InIntSlice(int(data.Status), statusSlice) {
		return errors.New("状态参数错误")
	}
	// 检查用户状态是否和要更新的状态一致
	user, err := dao.SysUser.FindOneById(data.Id, false)
	if err != nil {
		return errors.New("未查询到该用户")
	}
	if user.Status == data.Status {
		return errors.New("该用户状态已经发生变更，请重试")
	}
	user.Status = data.Status
	return dao.SysUser.Update(user)
}

func (u *SysUserService) CheckPhoneExists(phone string) bool {
	var num int64
	global.DB.Model(&entity.SysUser{}).Where("phone = ?", phone).Count(&num)
	return num > 0
}
