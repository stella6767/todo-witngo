package repository

import (
	"context"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"todo-app/.gen/postgres/public/model"
	"todo-app/.gen/postgres/public/table"
)

// 공개되는 인터페이스
type TodoRepository interface {
	CreateTodo(ctx context.Context, title string) (*model.Todo, error)
	GetTodos(ctx context.Context) ([]model.Todo, error)
	UpdateTodoStatus(ctx context.Context, id int32, completed bool) error
}

type todoRepository struct {
	//dbquery의 포인터
	db qrm.DB
}

// 생성자 함수
func NewTodoRepository(q qrm.DB) TodoRepository {
	return &todoRepository{db: q}
}

func (r *todoRepository) CreateTodo(ctx context.Context, title string) (*model.Todo, error) {
	stmt := table.Todo.INSERT(table.Todo.Title).
		VALUES(title).
		RETURNING(table.Todo.AllColumns)
	var todo model.Todo
	err := stmt.QueryContext(ctx, r.db, &todo)
	return &todo, err
}

func (r *todoRepository) GetTodos(ctx context.Context) ([]model.Todo, error) {
	stmt := table.Todo.SELECT(table.Todo.AllColumns)
	var todos []model.Todo
	err := stmt.QueryContext(ctx, r.db, &todos)
	return todos, err
}

func (r *todoRepository) UpdateTodoStatus(ctx context.Context, id int32, isComplete bool) error {
	_, err := table.Todo.UPDATE(table.Todo.Completed).
		SET(isComplete).
		WHERE(table.Todo.ID.EQ(postgres.Int32(id))).
		Exec(r.db)
	if err != nil {
		return err
	}
	return nil
}
