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

// 从cookie,header,query,postform上获取鉴权信息
func getParam(c *gin.Context, key string) (val string) {
	if val, _ = c.Cookie(key); val != "" {
		return
	} else if val = c.GetHeader(key); val != "" {
		return
	} else if val = c.Query(key); val != "" {
		return
	} else if val = c.PostForm(key); val != "" {
		return
	}
	return
}

// 鉴权
// JWT方式
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, user := "", ""
		if token = getParam(c, "token"); token == "" {
			response.FailError(c, apierr.UnAuthorized)
			c.Abort()
			return
		}
		if user = getParam(c, "user_id"); user == "" {
			response.FailError(c, apierr.UnAuthorized)
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
