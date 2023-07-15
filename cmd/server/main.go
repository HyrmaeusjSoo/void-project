package main

import (
	"fmt"
	"log"
	"void-project/global"
	"void-project/initialize"
	"void-project/internal/router"
)

func main() {
	defer func() {
		fmt.Println("")
		fmt.Println("按[回车]键退出...")
		fmt.Scanln()
	}()

	// 初始化配置
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库连接
	initialize.InitRepository()

	// 初始化Server
	r := initialize.InitServer()
	// 绑定路由
	router.SetApiRouter(r) // api router
	router.SetWebRouter(r) // view router (html templates)

	// 启动监听服务
	err := r.Run(global.Config.System.ListenAddr)
	if err != nil {
		log.Fatal(err)
	}
}
