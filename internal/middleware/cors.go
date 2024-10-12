package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 允许跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Origin, X-Requested-With, X-CSRF-Token, Accept, Accept-Encoding, Authorization, AccessToken, Token, userid, user_id, token")
		c.Header("Access-Control-Allow-Credentials", "true")
		// c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Origin, Access-Control-Allow-Headers, Access-Control-Max-Age, Content-Type, Content-Length")
		c.Header("Access-Control-Max-Age", "172800")
		// c.Set("Content-Type", "application/json")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
