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
}

// 实现error接口，return FailError时参数和error通用
func (jr JsonResponse) Error() string {
	return jr.Message
}

// 创建JsonResponse的错误结构体
func NewJsonResError(code int, message string) JsonResponse {
	return JsonResponse{
		Code:    code,
		Message: message,
	}
}

// 分页对象
type JsonResponsePage struct {
	List  any `json:"list,omitempty"`
	Total int `json:"total"`
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

// 返回定义错误
func FailError(c *gin.Context, err error) {
	if jr, ok := err.(JsonResponse); ok {
		c.JSON(jr.Code/100000, jr)
	} else {
		c.JSON(http.StatusOK, JsonResponse{
			Code:    1,
			Message: err.Error(),
		})
	}
}
