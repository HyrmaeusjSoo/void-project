package middleware

import (
	"strconv"
	"time"
	"void-project/internal/api/response"
	"void-project/internal/api/response/apierr"
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
			response.FailError(c, apierr.Unauthorized)
			c.Abort()
			return
		}

		if user, _ = c.Cookie("userId"); user != "" {
		} else if user = c.GetHeader("userId"); user != "" {
		} else if user = c.Query("userId"); user != "" {
		} else if user = c.PostForm("userId"); user != "" {
		} else {
			response.FailError(c, apierr.Unauthorized)
			c.Abort()
			return
		}
		userId, err := strconv.Atoi(user)
		if err != nil {
			response.FailError(c, apierr.AuthInvalidUserId)
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.FailError(c, apierr.AuthInvalidToken)
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			response.FailError(c, apierr.AuthExpired)
			c.Abort()
			return
		}
		if claims.UserID != uint(userId) {
			response.FailError(c, apierr.AuthUserIdMismatch)
			c.Abort()
			return
		}
		c.Set("userId", userId)

		c.Next()
	}
}
