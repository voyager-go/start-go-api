package service

import (
	"errors"
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

func (u *SysUserService) CheckPhoneExists(phone string) bool {
	var num int64
	global.DB.Model(&entity.SysUser{}).Where("phone = ?", phone).Count(&num)
	return num > 0
}
