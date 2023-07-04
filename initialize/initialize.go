package initialize

import (
	"io"
	"os"
	"void-project/global"
	"void-project/internal/repository/driver"
	"void-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

// 初始化配置文件
func InitConfig() {
	global.InitConfig()
}

// 初始化数据库连接
func InitRepository() {
	driver.InitMySQL()
	driver.InitRedis()
	driver.InitSQLite()
}

// 初始化Server日志
func InitServerLog() func() {
	file, err := logger.OpenLogFile(logger.ServerLevel)
	if err != nil {
		panic(err)
	}
	// defer file.Close()
	gin.DefaultWriter = io.MultiWriter(os.Stdout, file)
	gin.ForceConsoleColor()
	return func() {
		file.Close()
	}
}
