package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
	"void-project/global"
	"void-project/initialize"
	"void-project/internal/router"
	"void-project/pkg/logger"
)

// Web服务主程序
// 首先要初始化各种配置。可以在main入口，也可以都放在./initialize/下进行，从而让main方法更清爽。
// 再启动Gin服务。给Gin绑定自定义日志的io.Writer，给Gin绑定路由。
// 如果放在服务器上做成守护进程的话，要把fmt.Scanln类似的卡住控制台的代码都去掉。开发阶段可以留着看服务启动的最后报错信息。
func main() {
	// 意外panic捕捉
	defer func() {
		if err := recover(); err != nil {
			logger.LogServer(err)
			logger.LogServer(string(debug.Stack()))
		}
		logger.LogServer("服务已停止")
		fmt.Println("按[回车]键退出...")
		fmt.Scanln()
	}()

	// 初始化配置
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库连接
	initialize.InitRepository()
	// 初始化Jwt
	initialize.InitAuth()
	// 初始化翻译接口配置
	initialize.InitTranslate()

	// 初始化Server
	r := initialize.InitServer()
	// 绑定路由
	router.SetApiRouter(r) // api router
	router.SetWebRouter(r) // view router (html templates)

	// 性能分析 http://address:port/debug/pprof/
	if global.Configs.System.PProf {
		go func() {
			logger.LogServer(http.ListenAndServe(global.Configs.System.PProfAddr, nil))
		}()
	}

	// 启动监听服务
	err := r.Run(global.Configs.System.ListenAddr)
	if err != nil {
		logger.LogError("服务启动失败或意外关闭：" + err.Error())
	}
}
