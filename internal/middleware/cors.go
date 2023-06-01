package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Max-Age", "172800")
		// c.Header("Access-Control-Expose-Headers", "*")
		// c.Header("Access-Control-Allow-Credentials", "true")
		// c.Set("content-type", "application/json")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			// c.JSON(http.StatusOK, "response OK")
		}

		c.Next()
	}
}
