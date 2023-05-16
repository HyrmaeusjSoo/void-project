package initialize

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"chat/internal/repository/db"
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
	_, currentFile, _, _ := runtime.Caller(0)
	file, _ := os.Open(path.Dir(path.Dir(currentFile)) + "/config/mysql.json")
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&op); err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", op.User, op.Password, op.Host, op.Port, op.DBName)

	//sql语句记录
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer 日志输出目标，前缀和日志包含的内容
		logger.Config{
			SlowThreshold:             time.Second, // 慢SQL阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略‘记录未找到’错误
			Colorful:                  true,        // 彩色打印
		},
	)

	var err error
	db.MySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

// 初始化SQLite连接
func InitSQLite() {
	dir, _ := os.Getwd()
	//读取配置文件
	op := struct {
		Path     string
		User     string
		Password string
	}{}
	_, currentFile, _, _ := runtime.Caller(0)
	file, _ := os.Open(path.Dir(path.Dir(currentFile)) + "/config/sqlite.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err2 := decoder.Decode(&op)
	if err2 != nil {
		panic(err2)
	}

	var err error
	db.SQLite, err = gorm.Open(sqlite.Open(dir+op.Path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
