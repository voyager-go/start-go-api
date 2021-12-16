package main

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/bootstrap"
)

func main() {
	bootstrap.RunService()
	// result, err := bootstrap.Redis.Ping(context.Background()).Result()
	r := gin.Default()
	r.Run(":8090")
}
