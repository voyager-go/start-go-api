package main

import (
	"github.com/voyager-go/start-go-api/bootstrap"
	"github.com/voyager-go/start-go-api/config"
	"github.com/voyager-go/start-go-api/pkg/validator"
	"github.com/voyager-go/start-go-api/router"
)

func main() {
	// 程序启动时需要加载的服务
	bootstrap.BootService()
	// 引入验证翻译器
	validator.NewValidate()
	// 注册路由
	r := router.Register()
	// 程序启动
	r.Run(":" + config.Conf.Server.Port)
}
