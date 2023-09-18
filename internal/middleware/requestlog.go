package middleware

import (
	"void-project/pkg/logger/slog"

	"github.com/gin-gonic/gin"
)

func WriteRequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Info("request", "ip", c.ClientIP(), "method", c.Request.Method, "url", c.Request.URL.Path)
		c.Next()
	}
}
