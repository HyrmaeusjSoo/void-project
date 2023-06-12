package router

import (
	"chat/internal/middleware"
	"chat/internal/view"
	"chat/pkg"

	"github.com/gin-gonic/gin"
)

var vr = view.ViewHandler{}

func SetWebRouter(gin *gin.Engine) {
	gin.LoadHTMLGlob(pkg.GetRootPath() + "/web/template/*")
	gin.Static("static", pkg.GetRootPath()+"/web/static/")
	v1 := gin.Group("web/v1")

	// v1.Use(middleware.Cors())
	{
		v1.GET("", vr.Index)
		v1.GET("/chat", middleware.JWTAuth(), vr.Chat)
	}

}
