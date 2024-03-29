package middleware

import (
	"errors"
	"strconv"
	"time"
	"void-project/internal/api/response"
	"void-project/internal/api/response/apierr"
	"void-project/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// 鉴权
// JWT方式
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

		if user, _ = c.Cookie("user_id"); user != "" {
		} else if user = c.GetHeader("user_id"); user != "" {
		} else if user = c.Query("user_id"); user != "" {
		} else if user = c.PostForm("user_id"); user != "" {
		} else {
			response.FailError(c, apierr.Unauthorized)
			c.Abort()
			return
		}
		userId, err := strconv.Atoi(user)
		if err != nil {
			response.FailError(c, apierr.AuthInvalidUserId, err)
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				response.FailError(c, apierr.AuthExpired)
			} else {
				response.FailError(c, apierr.AuthInvalidToken, err)
			}
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
