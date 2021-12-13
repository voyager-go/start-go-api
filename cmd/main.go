package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/config"
)

func main() {
	fmt.Println(config.Conf.Mode)
	r := gin.Default()
	r.Run(":8090")
}
