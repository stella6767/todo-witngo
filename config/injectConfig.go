package config

import (
	"database/sql"
	"todo-app/internal/handler"
	"todo-app/internal/repository"
	"todo-app/internal/service"
)

func InitAppDependency(db *sql.DB) *handler.TodoHandler {

	// 레이어 초기화
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	return todoHandler
}
