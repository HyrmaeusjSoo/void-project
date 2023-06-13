package router

import (
	"chat/internal/api/handler"
	"chat/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(gin *gin.Engine) {
	var (
		userApi      = handler.NewUser()
		astroDictApi = handler.NewAstroDict()
		messageApi   = handler.NewMessage()
	)

	gin.Use(middleware.Cors())
	v1 := gin.Group("api/v1")
	// v1.Use(middleware.Cors())

	u := v1.Group("user")
	{
		u.POST("", userApi.Register)
		u.POST("login", userApi.Login)
		u.GET("/:id", middleware.JWTAuth(), userApi.Fetch)
		u.GET("", middleware.JWTAuth(), userApi.List)
		u.PUT("/:id", middleware.JWTAuth(), userApi.Update)
		u.DELETE("/:id", middleware.JWTAuth(), userApi.Delete)
		u.GET("online/:selfid", middleware.JWTAuth(), messageApi.OnLine)
	}

	ad := v1.Group("astro").Use(middleware.JWTAuth())
	{
		ad.GET("/:name", astroDictApi.Fetch)
	}

	/* rel := v1.Group("relation").Use(middleware.JWTAuth())
	{
		rel.GET("/:id", nil)
		rel.GET("", nil)
		rel.POST("", nil)
		rel.PUT("/:id", nil)
		rel.DELETE("/:id", nil)
	} */

	msg := v1.Group("message").Use(middleware.JWTAuth())
	{
		msg.GET("/send", messageApi.SendUserMsg)
	}
}
