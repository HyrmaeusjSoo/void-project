package driver

import (
	"fmt"
	"time"
	"void-project/global"
	"void-project/pkg"
	log "void-project/pkg/logger"

	"github.com/glebarez/sqlite" //gorm.io/driver/sqlite
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	MySQL     *gorm.DB
	SQLite    *gorm.DB
	SQLServer *gorm.DB
	Redis     *redis.Client
)

// 初始化MySQL数据库连接
func InitMySQL() {
	// 读取配置文件
	op := global.Config.DB.MySQL
	isColorful := pkg.IfElse(global.Config.System.Mode == "release", false, true)

	var err error
	MySQL, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", op.User, op.Password, op.Host, op.Port, op.DBName)), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		// SQL语句记录日志
		Logger: logger.New(
			log.NewSQLLogger(),
			logger.Config{
				SlowThreshold:             time.Second, // 慢SQL阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略‘记录未找到’错误
				Colorful:                  isColorful,  // 彩色打印
			},
		),
	})
	if err != nil {
		panic(err)
	}
}

// 初始化SQLite连接
// ！这里sqlite驱动没使用gorm下的，因为它使用了cgo，这样的话交叉编译时候底层依赖库、编译工具链都不一样会导致编译时报错
func InitSQLite() {
	//读取配置文件
	op := global.Config.DB.SQLite
	isColorful := pkg.IfElse(global.Config.System.Mode == "release", false, true)

	var err error
	SQLite, err = gorm.Open(sqlite.Open(pkg.GetRootPath()+op.Path), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		// SQL语句记录
		Logger: logger.New(
			log.NewSQLLogger(),
			logger.Config{
				SlowThreshold:             time.Second, // 慢SQL阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略‘记录未找到’错误
				Colorful:                  isColorful,  // 彩色打印
			},
		),
	})
	if err != nil {
		panic(err)
	}
}

// 初始化Redis连接
func InitRedis() {
	op := global.Config.Cache.Redis
	Redis = redis.NewClient(&op)
}
