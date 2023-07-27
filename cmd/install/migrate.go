package main

import (
	"fmt"
	"void-project/initialize"
	"void-project/internal/model"
	"void-project/internal/repository/driver"
)

func main() {
	defer func() {
		fmt.Println("")
		fmt.Println("按[回车]键退出...")
		fmt.Scanln()
	}()

	initialize.InitConfig()
	driver.InitMySQL()

	err := driver.MySQL.AutoMigrate(&model.User{}, &model.Message{})
	if err != nil {
		panic(err)
	}
}
