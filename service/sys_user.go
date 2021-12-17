package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/util/gconv"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/dao"
	"github.com/voyager-go/start-go-api/entity"
	"github.com/voyager-go/start-go-api/global"
	"github.com/voyager-go/start-go-api/pkg/auth"
	"github.com/voyager-go/start-go-api/pkg/util"
	"time"
)

type SysUserService struct{}

var User = SysUserService{}

// CreateUser 创建用户
func (u *SysUserService) CreateUser(r entity.SysUserServiceCreateReq) error {
	if dao.SysUser.CheckPhoneExists(r.Phone, 0) {
		return errors.New("该手机号已经存在")
	}
	statusSlice := []int{entity.SysUserStatusForbidden, entity.SysUserStatusNormal}
	if !util.InIntSlice(int(r.Status), statusSlice) {
		return errors.New("用户状态有误")
	}
	var user entity.SysUser
	err := gconv.Struct(r, &user)
	if err != nil {
		return errors.New("请求参数错误")
	}
	return dao.SysUser.Create(user)
}

// UpdateUser 更新用户信息
func (u *SysUserService) UpdateUser(r entity.SysUserServiceUpdateReq) error {
	if dao.SysUser.CheckPhoneExists(r.Phone, r.Id) {
		return errors.New("该手机号已经存在")
	}
	statusSlice := []int{entity.SysUserStatusForbidden, entity.SysUserStatusNormal}
	if !util.InIntSlice(int(r.Status), statusSlice) {
		return errors.New("用户状态有误")
	}
	user, err := dao.SysUser.FindOneById(r.Id, false)
	if err != nil || user.ID == 0 {
		return errors.New("未查询到该用户")
	}
	err = gconv.Struct(r, &user)
	if err != nil {
		return errors.New("请求参数错误")
	}
	return dao.SysUser.Update(user)
}

// Login 登录并发放令牌
func (u SysUserService) Login(r entity.SysUserServiceTokenReq) (string, error) {
	user, err := dao.SysUser.FindOneByPhone(r.Phone)
	if err != nil {
		return "", errors.New("未查找到该用户")
	}
	err = dao.SysUser.CheckUserPassword(&user, r.Password)
	if err != nil {
		return "", errors.New("手机号与密码不匹配")
	}
	jwtToken, err := auth.GenerateJwtToken(config.Conf.Server.JwtSecret, config.Conf.TokenExpire, user, config.Conf.Server.TokenIssuer)
	if err != nil {
		return "", errors.New("token生成失败")
	}
	marshal, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	err = global.Redis.Set(context.Background(), config.Conf.Redis.LoginPrefix+gconv.String(user.ID), string(marshal), time.Duration(config.Conf.Server.TokenExpire)*time.Second).Err()
	if err != nil {
		return "", errors.New("用户信息持久化失败")
	}
	return jwtToken, nil
}
