package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"todo-app/config"
	"todo-app/internal/db"
	"todo-app/internal/handler"
	"todo-app/internal/repository"
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

	// Gin 설정
	router := gin.Default()

	// 템플릿 등록
	router.GET("/", func(c *gin.Context) {
		todos, _ := todoRepo.GetTodos(c.Request.Context(), 1) // 임시 유저ID
		c.HTML(http.StatusOK, "", ui.Index(todos))
	})

	// API 라우트
	api := router.Group("/api/v1")
	{
		todoHandler := handler.NewTodoHandler(todoRepo)
		api.POST("/todos", todoHandler.CreateTodo)
	}

	// HTMX 에셋 제공
	router.Static("/static", "./static")

	log.Fatal(router.Run(":" + cfg.Port))
}
