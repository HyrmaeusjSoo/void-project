package middleware

import (
	"chat/pkg/jwt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		if token == "" {
			token = c.Request.Header.Get("token")
		}
		user := c.Query("userId")
		if user == "" {
			user = c.Request.Header.Get("userId")
		}
		userId, err := strconv.Atoi(user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "userId不正确",
			})
			c.Abort()
			return
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "请登录",
			})
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "token失效",
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "授权已过期",
			})
			c.Abort()
			return
		}
		if claims.UserID != uint(userId) {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "您的登录不合法",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
