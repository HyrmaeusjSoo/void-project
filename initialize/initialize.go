package initialize

import (
	"chat/internal/repository/driver"
)

func InitRepository() {
	driver.InitMySQL()
	driver.InitRedis()
	driver.InitSQLite()
}
