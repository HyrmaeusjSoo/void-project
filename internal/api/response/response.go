package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response返回Json对象
type JsonResponse struct {
	Code    int    `json:"code"`           // 状态码
	Message string `json:"message"`        // 消息内容
	Data    any    `json:"data,omitempty"` // 返回结构数据
	// HttpCode int    // http状态码
}

// 返回成功
func Success(ctx *gin.Context, data ...any) {
	ctx.JSON(http.StatusOK, JsonResponse{
		Code:    http.StatusOK,
		Message: "请求成功",
		Data:    data,
	})
}

// 返回失败
func Fail(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, JsonResponse{
		Code:    code,
		Message: msg,
	})
}
