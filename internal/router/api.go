package router

import (
	"void-project/internal/api/handler"
	"void-project/internal/middleware"

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
		u.POST("avatar", middleware.JWTAuth(), userApi.Avatar)
	}

	ad := v1.Group("astro").Use(middleware.JWTAuth())
	{
		ad.GET("/:name", astroDictApi.Fetch)
	}

	msg := v1.Group("message").Use(middleware.JWTAuth())
	{
		msg.GET("/send", messageApi.SendUserMsg)
		msg.GET("/online", messageApi.OnLine)
		msg.GET("", messageApi.List)
	}
}
