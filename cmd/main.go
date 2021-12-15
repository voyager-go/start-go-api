package main

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/bootstrap"
	_ "github.com/voyager-go/start-go-api/bootstrap"
)

func main() {
	bootstrap.RunService()
	r := gin.Default()
	r.Run(":8090")
}
