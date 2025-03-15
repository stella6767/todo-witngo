package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"todo-app/config"
	"todo-app/internal/db"
	"todo-app/internal/repository"
	"todo-app/internal/router"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// DB 연결
	pool, err := pgxpool.New(context.Background(), cfg.DBURL)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer pool.Close()

	// Repository 초기화
	queries := db.New(pool)
	todoRepo := repository.NewTodoRepository(queries)

	router.Router(todoRepo, cfg)

}
