package main

import (
	"database/sql"
	"embed"
	"github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"todo-app/config"
	"todo-app/internal/router"
)

//go:embed assets/*
var StaticAssets embed.FS

// main문이 실행되기 전에 실행
func init() {
	profile := os.Getenv("GO_PROFILE")
	config.LoadConfig(profile)
}

func main() {

	db := config.LoadDB()

	todoHandler := config.InitAppDependency(db)
	r := router.NewRouter(todoHandler)

	// 임베디드 파일 시스템 설정
	assets, _ := fs.Sub(StaticAssets, "assets")
	r.StaticFS("/assets", http.FS(assets))

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
