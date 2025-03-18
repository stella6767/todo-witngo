package service

import (
	"context"
	"todo-app/.gen/postgres/public/model"
	"todo-app/internal/dto"
	errUtil "todo-app/internal/errUtils"
	"todo-app/internal/repository"
)

type TodoService struct {
	repo      repository.TodoRepository
	txHandler repository.TransactionHandler
}

func NewTodoService(repo repository.TodoRepository, handler *repository.TransactionHandler) *TodoService {
	return &TodoService{repo: repo, txHandler: *handler}
}

func (s *TodoService) CreateTodo(ctx context.Context, task string) (*model.Todo, error) {

	var todo *model.Todo

	err := s.txHandler.Execute(ctx, func(ctx context.Context) error {
		result, err := s.repo.CreateTodo(ctx, task)
		if err != nil {
			return errUtil.Wrap(err)
		}
		todo = result
		return nil
	})

	if err != nil {
		return nil, errUtil.Wrap(err)
	}

	return todo, nil
}

func (s *TodoService) GetTodos(ctx context.Context) ([]model.Todo, error) {
	return s.repo.GetTodos(ctx)
}

func (s *TodoService) GetTodosByPage(ctx context.Context, pageable dto.Pageable) (dto.PageResult[model.Todo], error) {

	var todos dto.PageResult[model.Todo]

	err := s.txHandler.Execute(ctx, func(ctx context.Context) error {
		result, err := s.repo.GetTodosByPage(ctx, pageable)
		if err != nil {
			return errUtil.Wrap(err)
		}
		todos = result
		return nil
	})

	if err != nil {
		return dto.PageResult[model.Todo]{}, errUtil.Wrap(err)
	}

	return todos, nil
}

func (s *TodoService) UpdateTodoStatus(ctx context.Context, id int32) (model.Todo, error) {
	var todo model.Todo
	err := s.txHandler.Execute(ctx, func(ctx context.Context) error {
		result, err := s.repo.UpdateTodoStatus(ctx, id)
		if err != nil {
			return errUtil.Wrap(err)
		}
		todo = result
		return nil
	})
	if err != nil {
		return model.Todo{}, errUtil.Wrap(err)
	}
	return todo, nil
}

func (s *TodoService) DeleteTodoById(ctx context.Context, id int) error {
	return s.repo.DeleteTodoById(ctx, id)
}
