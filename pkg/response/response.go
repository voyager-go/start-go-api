package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JsonResponse 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    Code        `json:"code"`    // 错误码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// Response 基础返回
func Response(ctx *gin.Context, code Code, message string, data interface{}) {
	ctx.JSON(http.StatusOK, JsonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// Ok 正确返回
func Ok(ctx *gin.Context) {
	Response(ctx, Success, CodeMap[Success], interface{}(nil))
}

// OkWithMessage 自定义信息的正确返回
func OkWithMessage(ctx *gin.Context, message string) {
	Response(ctx, Success, message, interface{}(nil))
}

// OkWithData 携带数据信息的正确返回
func OkWithData(ctx *gin.Context, Data interface{}) {
	Response(ctx, Success, CodeMap[Success], Data)
}

// OkWithDetail 自定义的正确返回
func OkWithDetail(ctx *gin.Context, message string, Data interface{}) {
	Response(ctx, Success, message, Data)
}

// Fail 失败返回
func Fail(ctx *gin.Context) {
	Response(ctx, Failed, CodeMap[Failed], interface{}(nil))
}

// FailWithMessage 携带信息的失败返回
func FailWithMessage(ctx *gin.Context, message string) {
	Response(ctx, Failed, message, interface{}(nil))
}

// FailWithDetail 使用内置的错误枚举返回
func FailWithDetail(ctx *gin.Context, code Code) {
	Response(ctx, code, CodeMap[code], interface{}(nil))
}
