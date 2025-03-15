package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"todo-app/internal/db"
)

// 공개되는 인터페이스
type TodoRepository interface {
	CreateTodo(ctx context.Context, userID int32, title string) (db.Todo, error)
	GetTodos(ctx context.Context, userID int32) ([]db.Todo, error)
	UpdateTodoStatus(ctx context.Context, id int32, completed bool) error
}

type todoRepository struct {
	//dbquery의 포인터
	queries *db.Queries
}

// 생성자 함수
func NewTodoRepository(q *db.Queries) TodoRepository {
	return &todoRepository{queries: q}
}

func test(todo db.Todo) {

	//todo.ID.String()

}

func (r *todoRepository) CreateTodo(ctx context.Context, userID int32, title string) (db.Todo, error) {
	return r.queries.CreateTodo(ctx, db.CreateTodoParams{
		UserID: pgtype.Int4{Int32: userID, Valid: true},
		Title:  title,
	})
}

func (r *todoRepository) GetTodos(ctx context.Context, userID int32) ([]db.Todo, error) {
	return r.queries.GetTodosByUser(ctx, pgtype.Int4{Int32: userID, Valid: true})
}

func (r *todoRepository) UpdateTodoStatus(ctx context.Context, userID int32, isComplete bool) error {
	return r.queries.UpdateTodoStatus(ctx, db.UpdateTodoStatusParams{
		ID:        userID,
		Completed: pgtype.Bool{Bool: isComplete},
	})
}
