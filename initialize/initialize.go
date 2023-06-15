package initialize

import (
	"void-project/internal/repository/driver"
)

func InitRepository() {
	driver.InitMySQL()
	driver.InitRedis()
	driver.InitSQLite()
}
