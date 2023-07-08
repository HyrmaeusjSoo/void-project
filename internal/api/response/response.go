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

// 分页对象
type JsonResponsePage struct {
	List  any `json:"list,omitempty"`
	Total int `json:"total,omitempty"`
}

func NewJsonResponseError(code int, message string) JsonResponse {
	return JsonResponse{
		Code:    code,
		Message: message,
	}
}

// 返回状态结果
func SuccessOk(c *gin.Context) {
	c.JSON(http.StatusOK, JsonResponse{
		Code:    http.StatusOK,
		Message: "请求成功",
	})
}

// 返回对象结果
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, JsonResponse{
		Code:    http.StatusOK,
		Message: "请求成功",
		Data:    data,
	})
}

// 返回分页
func SuccessPage(c *gin.Context, data any, total int) {
	c.JSON(http.StatusOK, JsonResponse{
		Code:    http.StatusOK,
		Message: "请求成功",
		Data: JsonResponsePage{
			List:  data,
			Total: total,
		},
	})
}

// 返回失败
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(code, JsonResponse{
		Code:    code,
		Message: msg,
	})
}

// 返回已定义错误
func FailError(c *gin.Context, err JsonResponse) {
	c.JSON(err.Code/100000, err)
}
