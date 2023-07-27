package response

import (
	"net/http"
	"void-project/internal/model/base"
	"void-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Response返回Json对象
type JsonResponse struct {
	Code    int    `json:"code"`           // 状态码
	Message string `json:"message"`        // 消息内容
	Data    any    `json:"data,omitempty"` // 返回结构数据
}

// 创建Response对象
func NewResponse(code int, message string, data any) JsonResponse {
	return JsonResponse{code, message, data}
}

// 创建成功对象
func NewSuccess(data ...any) JsonResponse {
	res := NewResponse(http.StatusOK, "请求成功", nil)
	if data != nil {
		if len(data) > 1 {
			res.Data = data
		} else {
			res.Data = data[0]
		}
	}
	return res
}

// 创建错误对象
func NewJsonResError(code int, message string) JsonResponse {
	return NewResponse(code, message, nil)
}

// 实现error接口，return FailError时参数和error通用
func (jr JsonResponse) Error() string {
	return jr.Message
}

// 分页对象
type JsonResponsePage struct {
	List  any `json:"list,omitempty"`
	Total int `json:"total"`
}

// Cursor分页对象
type JsonResponseCursor struct {
	List any       `json:"list,omitempty"`
	Next base.Next `json:"next"`
}

// 返回状态结果
func SuccessOk(c *gin.Context) {
	c.JSON(http.StatusOK, NewSuccess())
}

// 返回对象结果
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, NewSuccess(data))
}

// 返回分页
func SuccessPage(c *gin.Context, data any, total int) {
	c.JSON(http.StatusOK, NewSuccess(JsonResponsePage{data, total}))
}

// 返回cursor方式分页
func SuccessCursor(c *gin.Context, data any, next base.Next) {
	c.JSON(http.StatusOK, NewSuccess(JsonResponseCursor{data, next}))
}

// 返回失败
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(code, JsonResponse{Code: code, Message: msg})
}

// 返回定义错误
func FailError(c *gin.Context, err error) {
	if jr, ok := err.(JsonResponse); ok {
		c.JSON(jr.Code/100000, jr)
	} else {
		logger.LogError(err)
		c.JSON(http.StatusOK, NewJsonResError(1, err.Error()))
	}
}
