package driver

import (
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
	MySQL  *gorm.DB
	SQLite *gorm.DB
	//SQLServer *gorm.DB
	Redis *redis.Client
)

type DBType string

var (
	DBTypeMySQL     DBType = "MySQL"
	DBTypeSQLite    DBType = "SQLite"
	DBTypeSQLServer DBType = "SQLServer"
)

// 初始化数据库
func InitDatabase(typ DBType) (gormDB *gorm.DB) {
	// 读取配置文件
	var op global.Databaser
	var dialector gorm.Dialector
	switch typ {
	case DBTypeMySQL:
		op = &global.Configs.DB.MySQL
		dialector = mysql.Open(op.Dsn())
	case DBTypeSQLite:
		op = &global.Configs.DB.SQLite
		dialector = sqlite.Open(op.Dsn())
	}
	isColorful := pkg.IfElse(global.Configs.System.Mode == global.ReleaseMode, false, true)

	var err error
	gormDB, err = gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		// SQL语句记录日志
		Logger: logger.New(
			log.NewSQLLogger(),
			logger.Config{
				SlowThreshold:             time.Second,                    // 慢SQL阈值
				LogLevel:                  gormLogLevel(op.SQLLogLevel()), // 日志级别
				IgnoreRecordNotFoundError: true,                           // 忽略‘记录未找到’错误
				Colorful:                  isColorful,                     // 彩色打印
			},
		),
	})
	if err != nil {
		panic(err)
	}
	db, err := gormDB.DB()
	if err != nil {
		panic(err)
	}
	maxIdleConns, maxOpenConns, maxLifetime := op.ConnsNumber()
	if maxIdleConns > 0 {
		db.SetMaxIdleConns(maxIdleConns)
	}
	if maxOpenConns > 0 {
		db.SetMaxOpenConns(maxOpenConns)
	}
	if maxLifetime > 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(maxLifetime))
	}
	return
}

// 初始化MySQL连接
func InitMySQL() {
	MySQL = InitDatabase(DBTypeMySQL)
}

// 初始化SQLite连接
// ！这里sqlite驱动没使用gorm下的，因为它使用了cgo，这样的话交叉编译时候底层依赖库、编译工具链都不一样会导致编译时报错
func InitSQLite() {
	SQLite = InitDatabase(DBTypeSQLite)
}

// 初始化Redis连接
func InitRedis() {
	op := global.Configs.Cache.Redis
	Redis = redis.NewClient(&op)
}

func gormLogLevel(level string) logger.LogLevel {
	return map[string]logger.LogLevel{
		"silent": logger.Silent,
		"error":  logger.Error,
		"warn":   logger.Warn,
		"info":   logger.Info,
	}[level]
}
