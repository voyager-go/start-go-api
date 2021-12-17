package dao

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/util/gconv"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/entity"
	"github.com/voyager-go/start-go-api/global"
	"golang.org/x/crypto/bcrypt"
)

type SysUserDao struct{}

var SysUser = new(SysUserDao)

// FindOneById 根据用户编号查找一条用户记录
func (u *SysUserDao) FindOneById(uid int64, ifNeedNormal bool) (entity.SysUser, error) {
	user := entity.SysUser{}
	// 先从redis中取数据
	jsonData, err := global.Redis.Get(context.Background(), config.Conf.Redis.LoginPrefix+gconv.String(uid)).Result()
	if err == nil {
		err := json.Unmarshal([]byte(jsonData), &user)
		if err == nil {
			return user, nil
		}
	}
	// redis中取数据异常，再从mysql中取数据
	result := global.DB.Model(&entity.SysUser{}).Where("id = ?", uid).Find(&user)
	if ifNeedNormal && user.Status == entity.SysUserStatusForbidden {
		return user, errors.New("用户状态异常")
	}
	return user, result.Error
}

// FindOneByPhone 根据用户手机号查找一条用户记录
func (u *SysUserDao) FindOneByPhone(phone string) (entity.SysUser, error) {
	user := entity.SysUser{}
	result := global.DB.Model(&entity.SysUser{}).Where("phone = ?", phone).Find(&user)
	return user, result.Error
}

// CheckUserPassword 匹配当前用户密码是否正确
func (u SysUserDao) CheckUserPassword(user *entity.SysUser, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// Create 创建用户记录
func (u *SysUserDao) Create(userData entity.SysUser) error {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userData.Password = string(fromPassword)
	result := global.DB.Create(&userData)
	return result.Error
}

// Update 更新用户信息
func (u *SysUserDao) Update(userData entity.SysUser) error {
	result := global.DB.Save(&userData)
	return result.Error
}

// CheckPhoneExists 检查手机号是否存在，ignoreId 参数传0则表示全表查找 传用户ID则表示忽略掉该用户
func (u *SysUserDao) CheckPhoneExists(phone string, ignoreId int64) bool {
	var num int64
	q := global.DB.Model(&entity.SysUser{})
	if ignoreId == 0 {
		q.Where("phone = ?", phone)
	} else {
		q.Where("phone = ? AND id <> ?", phone, ignoreId)
	}
	q.Count(&num)
	return num > 0
}
