package initialize

import (
	"void-project/global"
	"void-project/internal/repository/driver"
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
