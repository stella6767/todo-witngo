package main

import (
	"todo-app/config"
	"todo-app/internal/router"
)

func main() {

	todoHandler, _ := config.InitApp()

	r := router.NewRouter(todoHandler)

	err := r.Run(":8080")

	if err != nil {
		return
	}
}
