package middleware

import (
	"net/http"
	"strconv"
	"time"
	"void-project/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, user := "", ""
		if token, _ = c.Cookie("token"); token != "" {
		} else if token = c.GetHeader("token"); token != "" {
		} else if token = c.Query("token"); token != "" {
		} else if token = c.PostForm("token"); token != "" {
		} else {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "请登录",
			})
			c.Abort()
			return
		}

		if user, _ = c.Cookie("userId"); user != "" {
		} else if user = c.GetHeader("userId"); user != "" {
		} else if user = c.Query("userId"); user != "" {
		} else if user = c.PostForm("userId"); user != "" {
		} else {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "userId为空",
			})
			c.Abort()
			return
		}
		userId, err := strconv.Atoi(user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "userId不正确",
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
