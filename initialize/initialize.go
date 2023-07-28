package initialize

import (
	"fmt"
	"void-project/global"
	"void-project/internal/model/base"
	"void-project/internal/repository/driver"
	"void-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

func init() {
	echo()
}

// 输出标记
func echo() {
	mark := base.NewMark()
	for _, v := range mark {
		fmt.Println(v)
	}
}

// 初始化自定义日志
func InitLogger() {
	logger.InitLogger(global.Config.System.Mode)
}

// 初始化配置信息
func InitConfig() {
	global.InitConfig()
}

// 初始化数据库连接
func InitRepository() {
	driver.InitMySQL()
	driver.InitRedis()
	driver.InitSQLite()
}

// 初始化Server
func InitServer() *gin.Engine {
	if global.Config.System.Mode == "release" {
		gin.SetMode(gin.ReleaseMode) //发布模式
		gin.DisableConsoleColor()    //禁用彩色日志
	} else {
		gin.ForceConsoleColor() //彩色日志
	}
	//自定义日志
	gin.DefaultWriter = logger.NewServerLogger()

	// Gin引擎实例
	return gin.Default()
}
