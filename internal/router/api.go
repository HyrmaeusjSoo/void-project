package router

import (
	"void-project/internal/api/handler"
	"void-project/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RESTful风格的API路由
func SetApiRouter(gin *gin.Engine) {
	var (
		userApi      = handler.NewUser()
		astroDictApi = handler.NewAstroDict()
		messageApi   = handler.NewMessage()
		visitorApi   = handler.NewVisitor()
	)

	gin.Use(middleware.Cors())
	v1 := gin.Group("api/v1")
	// v1.Use(middleware.Cors())

	// 用户系列
	u := v1.Group("user")
	{
		u.POST("", userApi.Register)
		u.POST("login", userApi.Login)
		u.GET("/:id", middleware.JWTAuth(), userApi.Fetch)
		u.GET("", middleware.JWTAuth(), userApi.List)
		u.PUT("/:id", middleware.JWTAuth(), userApi.Update)
		u.DELETE("/:id", middleware.JWTAuth(), userApi.Delete)
		u.PUT("avatar", middleware.JWTAuth(), userApi.Avatar)
	}

	// 天文学词典系列
	ad := v1.Group("astro").Use(middleware.JWTAuth())
	{
		ad.GET("/:name", astroDictApi.Fetch)
		ad.GET("remote/:name", astroDictApi.FetchRemote)
		ad.POST("/:lang", astroDictApi.Sync)
		ad.GET("translate", astroDictApi.Translate)
	}

	// 聊天系列
	msg := v1.Group("message").Use(middleware.JWTAuth())
	{
		msg.GET("/send", messageApi.SendUserMsg)
		msg.GET("/online", messageApi.OnLine)
		msg.GET("", messageApi.List)
	}

	// 访客系列
	v := v1.Group("visitor").Use(middleware.JWTAuth())
	{
		v.GET("/ip/:ip", middleware.WriteRequestLog(), visitorApi.IP)
		v.GET("/log", visitorApi.FetchLog)
		v.GET("/stat", visitorApi.Stat)
	}

}
