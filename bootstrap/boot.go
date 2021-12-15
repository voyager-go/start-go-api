package bootstrap

import (
	"fmt"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/pkg/lib"
	"github.com/voyager-go/start-go-api/pkg/util"
	"gorm.io/gorm"
)

// 定义服务列表
const (
	LogService   = `Logger`
	RedisService = `Redis`
	MysqlService = `MySQL`
)

type bootServiceMap map[string]func() error

var (
	Mysql  *gorm.DB    // MySQL数据库
	Logger *lib.Logger // 日志
)

// BootedService 已经加载的服务
var BootedService []string

// serviceMap 程序启动时需要自动载入的服务
var serviceMap = bootServiceMap{
	LogService:   BootLogger,
	MysqlService: BootMysql,
}

// BootLogger 将配置载入日志服务
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

// BootMysql 将配置载入mysql服务
func BootMysql() error {
	if Mysql != nil {
		return nil
	}
	dbCfg := lib.DataBaseConfig{
		Host:     config.Conf.Host,
		Port:     config.Conf.Mysql.Port,
		User:     config.Conf.User,
		Password: config.Conf.Mysql.Password,
		DbName:   config.Conf.DbName,
	}
	_, err := lib.NewMysql(dbCfg)
	if err == nil {
		fmt.Println("程序载入MySQL服务成功!")
	}
	return err
}

// RunService 引导程序初始化，加载服务失败时会引发恐慌
// 日志服务默认加载，其它服务可选
func RunService(services ...string) {
	serviceMap[LogService] = BootLogger
	if Logger != nil {
		Logger.Infof("服务列表已加载完成!")
	}
	if len(services) == 0 {
		services = serviceMap.keys()
	}
	BootedService = make([]string, 0)
	for key, boot := range serviceMap {
		if util.InStringSlice(key, services) {
			if err := boot(); err != nil {
				panic("程序服务启动失败:" + err.Error())
			}
			BootedService = append(BootedService, key)
		}
	}
}

// keys 获取bootServiceMap实例中所有的键
func (s bootServiceMap) keys() []string {
	keys := make([]string, 0)
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
