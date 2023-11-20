package main

import (
	"fmt"
	"runtime/debug"
	"void-project/initialize"
	"void-project/internal/model"
	"void-project/internal/repository/driver"
)

// 迁移数据库结构、数据
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
			fmt.Println("")
			fmt.Println("按[回车]键退出...")
			fmt.Scanln()
		}
	}()

	initialize.InitConfig()
	driver.InitMySQL()

	err := driver.MySQL.AutoMigrate(&model.User{}, &model.Message{})
	if err != nil {
		panic(err)
	}
}
