package main

import (
	"database/sql"
	"github.com/sirupsen/logrus"
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

	db := config.LoadDB()
	todoHandler := config.InitAppDependency(db)
	r := router.NewRouter(todoHandler)
	port := strconv.Itoa(config.GlobalConfig.Server.Port)
	err := r.Run(":" + port)

	if err != nil {
		return
	}

	// main function 에 둬야됨
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logrus.WithError(err).Error("db close error")
		}
	}(db)

}
