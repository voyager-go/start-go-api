package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/voyager-go/start-go-api/pkg/lib"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB      // MySQL数据库
	Logger *lib.Logger   // 日志
	Redis  *redis.Client // Redis连接池
)
