package main

import (
	"log"
	"void-project/initialize"
	"void-project/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	initialize.InitConfig()
	// 初始化数据库连接
	initialize.InitRepository()

	// 初始化Server日志
	logClose := initialize.InitServerLog()
	defer logClose()

	// Server模式 debug/release
	// gin.SetMode(gin.ReleaseMode)

	// Gin引擎实例
	r := gin.Default()

	// 绑定路由
	router.SetApiRouter(r) // api router
	router.SetWebRouter(r) // view router (html templates)

	// 启动监听服务
	err := r.Run(":5555")
	if err != nil {
		log.Fatal(err)
	}
}
