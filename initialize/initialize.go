package initialize

import (
	"time"
	"void-project/global"
	"void-project/internal/middleware"
	"void-project/internal/model/base"
	"void-project/internal/repository/driver"
	"void-project/pkg"
	"void-project/pkg/jwt"
	"void-project/pkg/logger"
	"void-project/pkg/logger/slog"
	"void-project/pkg/necromancy"
	"void-project/pkg/translation"

	"github.com/gin-gonic/gin"
)

func init() {
	base.EchoMark()
}

// 初始化配置信息
func InitConfig() {
	global.InitConfig()
}

// 初始化自定义日志
func InitLogger() {
	logger.InitLogger(pkg.GetRootPath()+global.LogDir, global.Configs.System.Mode)
	slog.InitSLog(pkg.GetRootPath()+global.SLogDir, global.Configs.System.Mode)
}

// 初始化数据库连接
func InitRepository() {
	if necromancy.NotEmpty(global.Configs.DB.MySQL) {
		driver.InitMySQL()
	}
	if necromancy.NotEmpty(global.Configs.Cache) {
		driver.InitRedis()
	}
	if necromancy.NotEmpty(global.Configs.DB.SQLite) {
		driver.InitSQLite()
	}
}

// 初始化Auth
func InitAuth() {
	jwt.InitJwt(time.Hour*time.Duration(global.Configs.System.AuthTokenExpire), global.Configs.System.AuthJwtSecret)
}

// 初始化翻译接口配置
func InitTranslate() {
	translation.InitVolc(global.Configs.System.VolcAccessKey, global.Configs.System.VolcSecretKey)
}

// 初始化Server
func InitServer() *gin.Engine {
	if global.Configs.System.Mode == global.ReleaseMode {
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
