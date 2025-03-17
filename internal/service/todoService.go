package service

import (
	"context"
	"todo-app/.gen/postgres/public/model"
	"todo-app/internal/dto"
	"todo-app/internal/repository"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(ctx context.Context, task string) (*model.Todo, error) {

	return s.repo.CreateTodo(ctx, task)
}

func (s *TodoService) GetTodos(ctx context.Context) ([]model.Todo, error) {
	return s.repo.GetTodos(ctx)
}

func (s *TodoService) GetTodosByPage(ctx context.Context, pageable dto.Pageable) (dto.PageResult[model.Todo], error) {

	return s.repo.GetTodosByPage(ctx, pageable)
}

func (s *TodoService) UpdateTodoStatus(ctx context.Context, id int32) (model.Todo, error) {
	return s.repo.UpdateTodoStatus(ctx, id)
}

func (s *TodoService) DeleteTodoById(ctx context.Context, id int) error {
	return s.repo.DeleteTodoById(ctx, id)
}
