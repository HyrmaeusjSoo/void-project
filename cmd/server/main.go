package main

import (
	"fmt"
	"log"
	"void-project/global"
	"void-project/initialize"
	"void-project/internal/router"
)

// Web服务主程序
// 首先要初始化各种配置。可以在main入口，也可以都放在./initialize/下进行，从而让main方法更清爽。
// 再启动Gin服务。给Gin绑定自定义日志的io.Writer，给Gin绑定路由。
// 如果放在服务器上做成守护进程的话，要把fmt.Scanln类似的卡住控制台的代码都去掉。开发阶段可以留着看服务启动的最后报错信息。
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
