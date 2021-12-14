package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/voyager-go/start-go-api/bootstrap"
)

func main() {
	r := gin.Default()
	r.Run(":8090")
}
