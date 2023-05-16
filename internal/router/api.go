package router

import (
	"chat/internal/api/handler"
	"chat/internal/middleware"

	"github.com/gin-gonic/gin"
)

var (
	userApi handler.User
)

func SetApiRouter(gin *gin.Engine) {
	v1 := gin.Group("v1")
	v1.Use(middleware.Cors())

	u := v1.Group("user")
	{
		u.POST("", userApi.Register)
		u.POST("login", userApi.Login)
		u.GET("/:id", middleware.JWTAuth(), userApi.Fetch)
		u.PUT("", middleware.JWTAuth(), userApi.Update)
	}

	relation := v1.Group("relation").Use(middleware.JWTAuth())
	{
		relation.GET("fetch", nil)
	}
}
