package main

import (
	"chat/internal/model"
	"chat/internal/repository/driver"
)

func main() {
	driver.InitMySQL()

	err := driver.MySQL.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

}
