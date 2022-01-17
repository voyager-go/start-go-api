package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	adapter *gormadapter.Adapter
)

func init() {
	adapter, _ = gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/startgoapi", true)
}

func main() {
	r := InitRouter()
	r.Run(":9999")
}

func InitRouter() *gin.Engine {
	enforcer, _ := casbin.NewEnforcer("rbac_model.conf", adapter)
	enforcer.LoadPolicy()
	r := gin.Default()
	r.Use(Authorize(enforcer))
	r.POST("api/v1/add", AddPolicy)
	r.DELETE("api/v1/delete", RemovePolicy)
	r.GET("api/v1/list", ListPolicy)
	r.GET("api/v1/show", ShowSample)
	return r
}

// AddPolicy 新增策略
func AddPolicy(ctx *gin.Context) {
	enforcer, _ := casbin.NewEnforcer("/Users/artist/Program/go/start-go-api/casbin/rbac_model.conf", adapter)
	enforcer.LoadPolicy()
	if ok, _ := enforcer.AddPolicy("root", "/api/v1/show", "GET"); !ok {
		ctx.JSON(http.StatusOK, gin.H{"message": "Policy Exists"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Add Policy Success"})
	}
	return
}

// RemovePolicy 删除策略
func RemovePolicy(ctx *gin.Context) {
	enforcer, _ := casbin.NewEnforcer("/Users/artist/Program/go/start-go-api/casbin/rbac_model.conf", adapter)
	enforcer.LoadPolicy()
	if ok, _ := enforcer.RemovePolicy("root", "/api/v1/show", "GET"); !ok {
		ctx.JSON(http.StatusOK, gin.H{"message": "Policy Not Exists"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Remove Policy Success"})
	}
	return
}

// ListPolicy 查询全部策略
func ListPolicy(ctx *gin.Context) {
	enforcer, _ := casbin.NewEnforcer("/Users/artist/Program/go/start-go-api/casbin/rbac_model.conf", adapter)
	enforcer.LoadPolicy()
	list := enforcer.GetPolicy()
	ctx.JSON(http.StatusOK, gin.H{"message": list})
}

// ShowSample 策略允许通行后可以访问到的方法
func ShowSample(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "receive request ..."})
}

func Authorize(e *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取当前请求的 URI
		obj := ctx.Request.URL.RequestURI()
		act := ctx.Request.Method
		// 获取用户的角色
		sub := "root"
		fmt.Println(sub, obj, act)
		// 判断策略中是否存在
		fmt.Println(e.Enforce(sub, obj, act))
		if ok, _ := e.Enforce(sub, obj, act); ok {
			ctx.JSON(http.StatusOK, gin.H{"message": "Authorize pass ..."})
			ctx.Next()
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Authorize forbidden ..."})
			ctx.Abort()
		}
	}
}
