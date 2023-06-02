package driver

import (
	"chat/pkg"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
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
	//读取配置文件
	op := struct {
		User     string
		Password string
		Host     string
		Port     uint16
		DBName   string
	}{}
	file, _ := os.Open(pkg.GetRootPath() + "/config/mysql.json")
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&op); err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", op.User, op.Password, op.Host, op.Port, op.DBName)

	var err error
	MySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// SQL语句记录
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer 日志输出目标，前缀和日志包含的内容
			logger.Config{
				SlowThreshold:             time.Second, // 慢SQL阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略‘记录未找到’错误
				Colorful:                  true,        // 彩色打印
			},
		),
	})
	if err != nil {
		panic(err)
	}
}

// 初始化SQLite连接
func InitSQLite() {
	//读取配置文件
	op := struct {
		Path     string
		User     string
		Password string
	}{}
	file, _ := os.Open(pkg.GetRootPath() + "/config/sqlite.json")
	defer file.Close()
	if err2 := json.NewDecoder(file).Decode(&op); err2 != nil {
		panic(err2)
	}

	var err error
	SQLite, err = gorm.Open(sqlite.Open(pkg.GetRootPath()+op.Path), &gorm.Config{
		// SQL语句记录
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer 日志输出目标，前缀和日志包含的内容
			logger.Config{
				SlowThreshold:             time.Second, // 慢SQL阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略‘记录未找到’错误
				Colorful:                  true,        // 彩色打印
			},
		),
	})
	if err != nil {
		panic(err)
	}
}

// 初始化Redis连接
func InitRedis() {
	op := redis.Options{}
	file, _ := os.Open(pkg.GetRootPath() + "/config/redis.json")
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&op); err != nil {
		panic(err)
	}
	Redis = redis.NewClient(&op)
}
