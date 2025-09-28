package middleware

import (
	"errors"
	"strconv"
	"time"
	"void-project/internal/view"
	"void-project/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// 检查是否登录
func CheckWebLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, user := "", ""
		if token = getParam(c, "token"); token == "" {
			view.ErrorPage(c) //apierr.UnLogin.Message
			c.Abort()
			return
		}
		if user = getParam(c, "user_id"); user == "" {
			view.ErrorPage(c) //apierr.UnLogin.Message
			c.Abort()
			return
		}
		userId, err := strconv.Atoi(user)
		if err != nil {
			view.ErrorPage(c) //apierr.AuthInvalidUserId.Message
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				view.ErrorPage(c) //apierr.AuthExpired.Message
			} else {
				view.ErrorPage(c) //apierr.AuthInvalidToken.Message
			}
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			view.ErrorPage(c) //apierr.AuthExpired.Message
			c.Abort()
			return
		}
		if claims.UserID != uint(userId) {
			view.ErrorPage(c) //apierr.AuthUserIdMismatch.Message
			c.Abort()
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
