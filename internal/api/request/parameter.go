package request

import (
	"strconv"
	"time"
	"void-project/internal/api/response"
	"void-project/internal/api/response/apierr"

	"github.com/gin-gonic/gin"
)

// 获取当前登录的userId
func GetAuthUserId(c *gin.Context) (userId uint) {
	switch id := c.MustGet("userId").(type) {
	case int:
		userId = uint(id)
	case uint:
		userId = id
	case string:
		intUid, _ := strconv.Atoi(id)
		userId = uint(intUid)
	default:
		userId = 0
	}
	return
}

// 获取Query("name")中的int类型参数值
func GetQueryInt(c *gin.Context, name string) int {
	res, _ := strconv.Atoi(c.Query(name))
	return res
}

// 获取Query("name")中的int类型参数值 包含错误信息
func GetQueryIntErr(c *gin.Context, name string) (int, error) {
	if c.Query(name) == "" {
		err := apierr.MissingRequiredParameter.Wrap(name)
		response.FailError(c, err)
		return 0, err
	}
	res, err := strconv.Atoi(c.Query(name))
	if err != nil {
		err = apierr.InvalidParameter.Wrap(name, err)
		response.FailError(c, err)
		return 0, err
	}
	return res, nil
}

// 获取Query("name")中的time.Time类型参数值
func GetQueryTime(c *gin.Context, name string) time.Time {
	t, _ := time.ParseInLocation(time.DateTime, c.Query(name), time.Local)
	return t
}

// 获取Param("name")中的int类型参数值
func GetParamInt(c *gin.Context, name string) int {
	res, _ := strconv.Atoi(c.Param(name))
	return res
}

// 获取Param("name")中的int类型参数值 包含错误信息
func GetParamIntErr(c *gin.Context, name string) (int, error) {
	if c.Param(name) == "" {
		err := apierr.MissingRequiredParameter.Wrap(name)
		response.FailError(c, err)
		return 0, err
	}
	res, err := strconv.Atoi(c.Param(name))
	if err != nil {
		err = apierr.InvalidParameter.Wrap(name, err)
		response.FailError(c, err)
		return 0, err
	}
	return res, nil
}
