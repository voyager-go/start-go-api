package lib

import (
	"github.com/go-redis/redis/v8"
)

// RdsConfig 配置信息
type RdsConfig struct {
	Addr     string
	Password string
	DbNum    int
}

// NewRedis 构造redis客户端
func NewRedis(cnf RdsConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cnf.Addr,
		Password: cnf.Password,
		DB:       cnf.DbNum,
	})
}
