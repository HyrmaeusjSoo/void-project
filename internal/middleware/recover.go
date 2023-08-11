package middleware

import (
	"runtime/debug"
	"void-project/internal/api/response"
	"void-project/internal/api/response/apierr"
	"void-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

// 拦截处理过程中的panic级别错误，记录日志并返回错误信息
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logger.LogError(err)
			logger.LogError(string(debug.Stack()))
			response.FailError(c, apierr.InternalServerError, err)
			c.Abort()
		}
	}()

	c.Next()
}
