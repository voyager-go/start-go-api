package bootstrap

import (
	"fmt"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/pkg/lib"
	"gorm.io/gorm"
)

// 定义服务列表
const (
	LogService   = `Logger`
	RedisService = `Redis`
	DbService    = `Db`
)

type bootServiceMap map[string]func() error

var (
	Db     *gorm.DB    // 数据库
	Logger *lib.Logger // 日志
)

// BootedService 已经加载的服务
var BootedService []string

// serviceMap 程序启动时需要自动载入的服务
var serviceMap = bootServiceMap{
	LogService: BootLogger,
}

func BootLogger() error {
	if Logger != nil {
		return nil
	}
	_, err := lib.NewLogger(config.Conf.DirPath, config.Conf.FileName)
	if err == nil {
		fmt.Println("程序载入日志服务成功! 模块为:" + config.Conf.FileName + ", 日志路径为:" + config.Conf.DirPath)
	}
	return err
}

func Init() {
	serviceMap[LogService] = BootLogger
	if Logger != nil {
		Logger.Infof("服务列表已加载完成!")
	}
}
