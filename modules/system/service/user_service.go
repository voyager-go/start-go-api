package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/global"
	"github.com/voyager-go/start-go-api/pkg/auth"
	"github.com/voyager-go/start-go-api/repositories"
	"strconv"
	"time"
)

type userService struct{}

var User = userService{}

// Create 创建一条记录
func (s *userService) Create(req *entities.UserServiceCreateReq) error {
	var user entities.User
	err := gconv.Struct(req, &user)
	if err != nil {
		return err
	}
	userRepository := repositories.NewUserRepository()
	if userRepository.FindOneByPhone(req.Phone).Error == nil {
		return errors.New("该手机号已经存在")
	}
	return userRepository.Create(&user).Error
}

// Update 更新用户信息
func (s *userService) Update(req entities.UserServiceUpdateReq) error {
	userRepository := repositories.NewUserRepository()
	datum := userRepository.FindOneById(req.ID)
	user, ok := datum.Result.(*entities.User)
	if datum.Error != nil || !ok {
		return errors.New("该用户不存在")
	}
	datum = userRepository.FindOneByPhone(req.Phone)
	userByPhone, ok := datum.Result.(*entities.User)
	if ok && datum.Error == nil {
		if userByPhone.ID != req.ID {
			return errors.New("该手机号已经存在")
		}
	}
	err := gconv.Struct(req, &user)
	if err != nil {
		return errors.New("请求参数错误")
	}
	return userRepository.Update(user).Error
}

// FindOne 根据用户ID查找用户信息
func (s *userService) FindOne(ID uint64) (*entities.User, error) {
	userRepository := repositories.NewUserRepository()
	datum := userRepository.FindOneById(ID)
	user, ok := datum.Result.(*entities.User)
	if datum.Error != nil || !ok {
		return nil, errors.New("该用户不存在")
	}
	return user, nil
}

// Login 用户登录并发放令牌
func (s *userService) Login(req entities.UserServiceTokenReq) (string, error) {
	userRepository := repositories.NewUserRepository()
	result := userRepository.FindOneByPhone(req.Phone)
	if result.Error != nil {
		return "", errors.New("未查找到该用户")
	}
	user := result.Result.(*entities.User)
	result = userRepository.CheckPassword(user.Password, req.Password)
	if result.Error != nil {
		return "", errors.New("手机号与密码不匹配")
	}
	jwtToken, err := auth.GenerateJwtToken(config.Conf.Server.JwtSecret, config.Conf.TokenExpire, user, config.Conf.Server.TokenIssuer)
	if err != nil {
		return "", errors.New("token生成失败")
	}
	roleResult := repositories.NewSysRoleUserRepository().FindAllByUserId(user.ID)
	if roleResult.Error != nil {
		return "", roleResult.Error
	}
	roleInfo := roleResult.Result.([]entities.SysRoleUser)
	if len(roleInfo) > 0 {
		for _, role := range roleInfo {
			user.RoleIds = append(user.RoleIds, role.RoleId)
		}
	}
	marshal, err := json.Marshal(user)
	fmt.Println(string(marshal))
	if err != nil {
		return "", errors.New("token编码失败")
	}
	err = global.Redis.Set(context.Background(), config.Conf.Redis.LoginPrefix+gconv.String(user.Model.ID), string(marshal), time.Duration(config.Conf.Server.TokenExpire)*time.Second).Err()
	if err != nil {
		return "", errors.New("用户信息持久化失败")
	}
	return jwtToken, nil
}

// Logout 退出登录
func (s *userService) Logout(ID int64) error {
	_, err := global.Redis.Del(context.Background(), config.Conf.Redis.LoginPrefix+strconv.FormatInt(ID, 10)).Result()
	if err != nil {
		return errors.New("会话过期，请重新登录")
	}
	return nil
}
