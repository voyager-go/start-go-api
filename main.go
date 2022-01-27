package main

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/urfave/cli/v2"
	"github.com/voyager-go/start-go-api/bootstrap"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/pkg/migrate"
	"github.com/voyager-go/start-go-api/pkg/validator"
	"github.com/voyager-go/start-go-api/router"
	schedule_server "github.com/voyager-go/start-go-api/schedule/server"
	schedule_worker "github.com/voyager-go/start-go-api/schedule/worker"
	"os"
	"runtime"
	"time"
)

var (
	// AppName 当前应用的名称
	AppName  = "start go api"
	AppUsage = "使用gin框架作为基础开发库，封装一套适用于面向api编程的快速开发结构"
	// AppPort 程序启动的端口
	AppPort string
	// BuildVersion 编译的app版本
	BuildVersion string
	// BuildAt 编译时间
	BuildAt string
	// openSchedule 是否开启任务调度服务
	openSchedule = true
	// taskServer 任务服务
	taskServer    *machinery.Server
	taskServerArg = "server"
	taskWorkerArg = "worker"
)

// stack 程序运行前的处理
func stack() *cli.App {
	buildInfo := fmt.Sprintf("%s-%s-%s-%s-%s", runtime.GOOS, runtime.GOARCH, BuildVersion, BuildAt, time.Now().Format("2006-01-02 15:04:05"))
	if openSchedule {
		var err error
		taskServer, err = schedule_server.InitMachineryServer()
		if err != nil {
			panic(err)
		}
	}
	return &cli.App{
		Name:    AppName,
		Version: buildInfo,
		Usage:   AppUsage,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "env",
				Value:       "dev",
				Usage:       "请选择配置文件 [dev | pre]",
				Destination: &config.ConfEnv,
			},
			&cli.StringFlag{
				Name:        "port",
				Value:       "8090",
				Usage:       "请选择启动端口",
				Destination: &AppPort,
			},
		},
		Action: func(context *cli.Context) error {
			// 初始化配置文件信息
			config.InitConfig()
			// 程序启动时需要加载的服务
			bootstrap.BootService()
			// 引入验证翻译器
			validator.NewValidate()
			// 自动生成表结构
			migrate.AuthMigrate()
			// 注册路由 启动程序
			return router.Register().Run(":" + AppPort)
		},
		Commands: []*cli.Command{
			{
				Name:  taskServerArg,
				Usage: "Launch application server",
				Action: func(context *cli.Context) error {
					if openSchedule {
						schedule_server.StartServer(taskServer)
					}
					return nil
				},
			},
			{
				Name:  taskWorkerArg,
				Usage: "Launch application worker",
				Action: func(context *cli.Context) error {
					if openSchedule {
						return schedule_worker.StartWorker(taskServer)
					}
					return nil
				},
			},
		},
	}
}

func main() {
	if err := stack().Run(os.Args); err != nil {
		panic(err)
	}
}
