package config

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"todo-app/internal/handler"
	"todo-app/internal/repository"
	"todo-app/internal/service"
)

func InitAppDependency(db *sql.DB) *handler.TodoHandler {

	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true, // 터미널 색상 강제 적용
		FullTimestamp: true, // 전체 타임스탬프 표시
	})

	// 레이어 초기화
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo, repository.NewTxHandler(db))
	todoHandler := handler.NewTodoHandler(todoService)

	return todoHandler
}
