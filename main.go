package main

import (
	"todo-app/config"
	"todo-app/internal/router"
)

// main문이 실행되기 전에 실행
func init() {
	//profile = os.Getenv("GO_PROFILE")
	config.LoadConfig()
}

func main() {

	todoHandler := config.InitAppDependency()
	
	r := router.NewRouter(todoHandler)
	err := r.Run(":8080")

	if err != nil {
		return
	}
}
