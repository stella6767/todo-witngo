package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"todo-app/config"
	"todo-app/internal/handler"
	"todo-app/internal/repository"
	"todo-app/internal/view"
)

func Router(
	todoRepo repository.TodoRepository,
	cfg config.Config,
) {
	// Gin 설정
	router := gin.Default()
	// 템플릿 등록
	router.GET("/", func(c *gin.Context) {
		//todos, _ := todoRepo.GetTodos(c.Request.Context(), 1) // 임시 유저ID
		c.HTML(http.StatusOK, "", view.Test())
	})

	// API 라우트
	api := router.Group(`/v1`)
	{
		todoHandler := handler.NewTodoHandler(todoRepo)
		api.POST("/todos", todoHandler.CreateTodo)
	}

	// HTMX 에셋 제공
	router.Static("/static", "./static")

	log.Fatal(router.Run(":" + cfg.Port))
}
