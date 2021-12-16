package api

import (
	"github.com/gin-gonic/gin"
	"github.com/voyager-go/start-go-api/pkg/response"
)

type CheckApi struct{}

var Check = CheckApi{}

func (c *CheckApi) Ping(ctx *gin.Context) {
	response.OkWithData(ctx, "pong!")
}
