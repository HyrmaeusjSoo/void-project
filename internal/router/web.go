package router

import (
	"void-project/internal/middleware"
	"void-project/internal/view"
	"void-project/pkg"

	"github.com/gin-gonic/gin"
)

// View路由，传统后端驱动的SSR方式
func SetWebRouter(gin *gin.Engine) {
	// 加载页面模板文件
	gin.LoadHTMLGlob(pkg.GetRootPath() + "/web/template/*")
	// 加载静态文件
	gin.Static("static", pkg.GetRootPath()+"/web/static")
	gin.Static("upload", pkg.GetRootPath()+"/web/upload")

	gin.Use(middleware.Cors())
	// 设置路由
	v1 := gin.Group("web")
	{
		vr := view.ViewHandler{}
		v1.GET("", vr.Index)
		v1.GET("index", vr.Index)
		v1.Use(middleware.CheckWebLogin())
		v1.GET("chat", vr.Chat)
		v1.GET("api", vr.Api)
		v1.GET("storage", vr.Storage)
		v1.GET("gallery", vr.Gallery)
	}

}
