package main

import (
	"chat/initialize"
	"chat/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	initialize.InitMySQL()

	r := gin.Default()
	router.SetApiRouter(r)
	err := r.Run(":5555")
	if err != nil {
		panic(err)
	}
}
