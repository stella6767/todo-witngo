package main

import (
	"strconv"
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

	port := strconv.Itoa(config.GlobalConfig.Server.Port)
	err := r.Run(":" + port)
	
	if err != nil {
		return
	}
}
