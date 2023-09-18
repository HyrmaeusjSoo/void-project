package initialize

import (
	"fmt"
	"void-project/global"
	"void-project/internal/middleware"
	"void-project/internal/model/base"
	"void-project/internal/repository/driver"
	"void-project/pkg"
	"void-project/pkg/logger"
	"void-project/pkg/logger/slog"

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

// 初始化配置信息
func InitConfig() {
	global.InitConfig()
}

// 初始化自定义日志
func InitLogger() {
	logger.InitLogger(pkg.GetRootPath()+"/runtime/log/", global.Config.System.Mode)
	slog.InitSLog(pkg.GetRootPath()+"/runtime/slog/", global.Config.System.Mode)
}

// 初始化数据库连接
func InitRepository() {
	driver.InitMySQL()
	driver.InitRedis()
	driver.InitSQLite()
}

// 初始化Server
func InitServer() *gin.Engine {
	if global.Config.System.Mode == global.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //发布模式
		gin.DisableConsoleColor()    //禁用彩色日志
	} else {
		gin.ForceConsoleColor() //彩色日志
	}
	// 自定义请求日志
	gin.DefaultWriter = logger.NewServerLogger()
	// 自定义错误日志
	// gin.DefaultErrorWriter = io.MultiWriter(logger.NewLogger(logger.ErrorLevel), os.Stdout)

	// Gin引擎实例
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger())

	// 自定义中间件替代gin.Recovery()，如果想使用默认的, 需要设置gin.DefaultErrorWriter
	r.Use(middleware.Recover)
	return r
}
