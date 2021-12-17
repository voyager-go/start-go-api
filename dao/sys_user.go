package dao

import (
	"errors"
	"github.com/voyager-go/start-go-api/entity"
	"github.com/voyager-go/start-go-api/global"
)

type SysUserDao struct{}

var SysUser = new(SysUserDao)

// FindOneById 根据用户编号查找一条用户记录
func (u *SysUserDao) FindOneById(uid int64, ifNeedNormal bool) (entity.SysUser, error) {
	user := entity.SysUser{}
	user.ID = uid
	result := global.DB.Find(&user)
	if ifNeedNormal && user.Status == entity.SysUserStatusForbidden {
		return user, errors.New("用户状态异常")
	}
	return user, result.Error
}

// Update 更新用户信息
func (u *SysUserDao) Update(userData entity.SysUser) error {
	result := global.DB.Save(userData)
	return result.Error
}
