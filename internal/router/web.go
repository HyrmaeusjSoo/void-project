package router

import (
	"chat/internal/middleware"
	"chat/internal/view"
	"chat/pkg"

	"github.com/gin-gonic/gin"
)

func SetWebRouter(gin *gin.Engine) {
	gin.LoadHTMLGlob(pkg.GetRootPath() + "/web/template/*")
	gin.Static("static", pkg.GetRootPath()+"/web/static/")

	v1 := gin.Group("web/")
	{
		vr := view.ViewHandler{}
		v1.GET("", vr.Index)
		v1.GET("/chat", middleware.JWTAuth(), vr.Chat)
	}

}
