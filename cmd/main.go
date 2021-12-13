package main

import (
	"github.com/gin-gonic/gin"
	"github.com/youeryuango/start-go-api/config"
)

func main() {
	//var c = config.Conf.Redis{}
	r := gin.Default()
	r.Run(":8090")
}
