package main

import (
	"github.com/voyager-go/start-go-api/bootstrap"
	"github.com/voyager-go/start-go-api/router"
)

func main() {
	bootstrap.RunService()
	//result, err := bootstrap.Redis.Ping(context.Background()).Result()
	r := router.Register()
	r.Run(":8090")
}
