package main

import (
	"void-project/initialize"
	"void-project/internal/model"
	"void-project/internal/repository/driver"
)

func main() {
	initialize.InitConfig()
	driver.InitMySQL()

	err := driver.MySQL.AutoMigrate(&model.User{}, &model.Message{})
	if err != nil {
		panic(err)
	}
}
