package main

import (
	"chat/initialize"
	"chat/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库连接
	initialize.InitRepository()

	// 启动监听服务，绑定路由等
	r := gin.Default()
	router.SetApiRouter(r)
	err := r.Run(":5555")
	if err != nil {
		panic(err)
	}
}
