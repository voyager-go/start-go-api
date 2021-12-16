package dao

import (
	"github.com/voyager-go/start-go-api/entity"
	"github.com/voyager-go/start-go-api/global"
)

type SysUserDao struct{}

var SysUser = new(SysUserDao)

// FindOneById 根据用户编号查找一条用户记录
func (u *SysUserDao) FindOneById(uid int64) (entity.SysUser, error) {
	user := entity.SysUser{}
	user.ID = uid
	result := global.DB.Find(&user)
	return user, result.Error
}
