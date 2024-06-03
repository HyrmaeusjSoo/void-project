package router

import (
	"void-project/internal/middleware"
	"void-project/internal/view"
	"void-project/pkg"

	"github.com/gin-gonic/gin"
)

// View路由，传统后端驱动的SSR方式
func SetWebRouter(gin *gin.Engine) {
	gin.LoadHTMLGlob(pkg.GetRootPath() + "/web/template/*")
	gin.Static("static", pkg.GetRootPath()+"/web/static/")
	gin.Static("upload", pkg.GetRootPath()+"/web/upload/")

	v1 := gin.Group("web")
	{
		vr := view.ViewHandler{}
		v1.GET("", vr.Index)
		v1.GET("index", vr.Index)
		v1.Use(middleware.JWTAuth())
		v1.GET("chat", vr.Chat)
		v1.GET("api", vr.Api)
		v1.GET("storage", vr.Storage)
	}

}
