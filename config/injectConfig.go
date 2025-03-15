package config

import (
	"todo-app/internal/handler"
	"todo-app/internal/repository"
	"todo-app/internal/service"
)

func InitApp() (*handler.TodoHandler, Person) {

	db := loadDB()
	// 레이어 초기화
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	person := Person{Name: "test"}

	return todoHandler, person
}

type Person struct {
	Name string
	Age  int
}
