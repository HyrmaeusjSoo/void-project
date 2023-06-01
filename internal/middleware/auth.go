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
		token := c.GetHeader("token")
		if token == "" {
			token = c.Query("token")
			if token == "" {
				token = c.PostForm("token")
			}
		}
		user := c.GetHeader("userId")
		if user == "" {
			user = c.Query("userId")
			if user == "" {
				user = c.PostForm("userId")
			}
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
				"message": "token已失效",
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
