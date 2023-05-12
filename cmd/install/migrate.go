package main

import (
	"chat/initialize"
	"chat/internal/model"
	"chat/internal/repository/db"
)

func main() {
	initialize.InitMySQL()

	err := db.MySQL.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

}
