package router

import (
	"chat/internal/api/handler"
	"chat/internal/middleware"

	"github.com/gin-gonic/gin"
)

var (
	userApi      handler.User
	astroDictApi handler.AstroDict
)

func SetApiRouter(gin *gin.Engine) {
	gin.Use(middleware.Cors())
	v1 := gin.Group("v1")
	// v1.Use(middleware.Cors())

	u := v1.Group("user")
	{
		u.POST("", userApi.Register)
		u.POST("login", userApi.Login)
		u.GET("/:id", middleware.JWTAuth(), userApi.Fetch)
		u.PUT("", middleware.JWTAuth(), userApi.Update)
	}

	ad := v1.Group("astro").Use(middleware.JWTAuth())
	{
		ad.GET("/:name", astroDictApi.Fetch)
	}

	relation := v1.Group("relation").Use(middleware.JWTAuth())
	{
		relation.GET("fetch", nil)
	}
}
