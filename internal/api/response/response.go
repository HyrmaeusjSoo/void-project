package response

import (
	"errors"
	"net/http"
	"strings"
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
func New(code int, message string, data any) JsonResponse {
	return JsonResponse{code, message, data}
}

// 创建成功对象
func NewSuccess(data ...any) JsonResponse {
	res := New(http.StatusOK, "请求成功", nil)
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
func NewError(code int, message string) JsonResponse {
	return New(code, message, nil)
}

// 实现标准库error接口，return FailError时自定义api错误与标准error兼容
func (jr JsonResponse) Error() string {
	return jr.Message
}

// 包装错误，额外的错误包装在JsonResponse内
func (jr JsonResponse) Wrap(errs ...any) JsonResponse {
	if len(errs) == 0 {
		return jr
	}
	jr.Message += join('，', errs...)
	return jr
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

// 返回状态
// 只返回成功状态，没有额外的数据信息
func SuccessOk(c *gin.Context) {
	c.JSON(http.StatusOK, NewSuccess())
}

// 返回数据
//
//	c => gin.Context上下文对象指针
//	data => 返回数据
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, NewSuccess(data))
}

// 返回分页
// total 要查询的记录总数
func SuccessPage(c *gin.Context, data any, total int) {
	Success(c, JsonResponsePage{data, total})
}

// 返回cursor方式分页
// next 查询下一页的游标
func SuccessCursor(c *gin.Context, data any, next base.Next) {
	Success(c, JsonResponseCursor{data, next})
}

// 返回失败
//
//	当返回自定义错误或者不使用内置的apierr.XXX类型JsonResponse时调用该方法
//
//	c => gin.Context上下文对象指针
//	code => 错误码，该错误码会同时用在HTTP.Status和JsonResponse.code
//	msg => 错误信息的描述
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(code, NewError(code, msg))
}

// 返回定义错误
//
//	c => gin.Context上下文对象指针
//	err => 错误，apierr下的JsonResponse错误或是标准库error
//	ext => 额外的错误信息或包装错误，标准库error接口的实现类型的或string错误信息
//
//	返回时进行多错误包装有两种方式：
//		1. 直接使用response.FailError(c, err, ext...)，额外错误拼接在JsonResponse.Message中返回；
//		2. 也先用apierr包装错误，在FailError时传入包装后的错误，额外错误也会拼接在JsonReponse.Message中返回：
//			2.1. err = apierr.XXX.Wrap(ext)
//			2.2. response.FailError(c, err)
func FailError(c *gin.Context, err error, ext ...any) {
	if jr, ok := err.(JsonResponse); ok {
		c.JSON(jr.Code/100000, jr.Wrap(ext...))
	} else if jr := func() JsonResponse { return JsonResponse{} }(); errors.As(err, &jr) {
		jr.Message = err.Error()
		c.JSON(jr.Code/100000, jr.Wrap(ext...))
	} else {
		logger.LogError(err)
		c.JSON(http.StatusBadRequest, NewError(40000000, err.Error()+join('，', ext...)))
	}
}

// 拼接错误
func join(sep rune, errs ...any) string {
	extBuilder := strings.Builder{}
	if len(errs) > 0 {
		for _, e := range errs {
			if s, ok := e.(string); ok {
				extBuilder.WriteRune(sep)
				extBuilder.WriteString(s)
			} else if err2, ok := e.(error); ok {
				extBuilder.WriteRune(sep)
				extBuilder.WriteString(err2.Error())
			}
		}
	}
	return extBuilder.String()
}
