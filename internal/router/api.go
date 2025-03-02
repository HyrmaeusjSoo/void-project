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
		storageApi   = handler.NewStorage()
	)

	gin.Use(middleware.Cors())
	v1 := gin.Group("api/v1")
	// v1.Use(middleware.Cors())

	// 用户系列
	u := v1.Group("user")
	{
		u.POST("", userApi.Register)
		u.POST("login", userApi.Login)
		u.Use(middleware.JWTAuth())
		u.GET("/:id", userApi.Fetch)
		u.GET("", userApi.List)
		u.PUT("", userApi.Update)
		u.DELETE("", userApi.Delete)
		u.PUT("avatar", userApi.Avatar)
		u.PUT("password", userApi.UpdatePassword)
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

	// 文件存储系列
	stg := v1.Group("storage").Use(middleware.JWTAuth())
	{
		stg.GET("", storageApi.List)
		stg.POST("", storageApi.Mkdir)
		stg.GET("/download", storageApi.Download)
		stg.POST("/upload", storageApi.Upload)
		stg.DELETE("", storageApi.Delete)
		stg.PUT("rename", storageApi.Rename)
		stg.PUT("copy", storageApi.Copy)
		stg.PUT("move", storageApi.Move)
	}

}
