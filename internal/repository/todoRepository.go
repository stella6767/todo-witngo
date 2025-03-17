package repository

import (
	"context"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/pkg/errors"
	"todo-app/.gen/postgres/public/model"
	"todo-app/.gen/postgres/public/table"
	"todo-app/internal/dto"
	errUtil "todo-app/internal/errUtils"
)

// 공개되는 인터페이스
type TodoRepository interface {
	CreateTodo(ctx context.Context, title string) (*model.Todo, error)
	GetTodos(ctx context.Context) ([]model.Todo, error)
	GetTodosByPage(ctx context.Context, pageable dto.Pageable) dto.PageResult[model.Todo]
	UpdateTodoStatus(ctx context.Context, id int32) (model.Todo, error)
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
	return &todo, errUtil.Wrap(err)
}

func (r *todoRepository) GetTodos(ctx context.Context) ([]model.Todo, error) {
	stmt := table.Todo.SELECT(table.Todo.AllColumns)
	var todos []model.Todo
	err := stmt.QueryContext(ctx, r.db, &todos)
	return todos, errUtil.Wrap(err)
}

func (r *todoRepository) GetTodosByPage(ctx context.Context, pageable dto.Pageable) dto.PageResult[model.Todo] {

	stmt := table.Todo.
		SELECT(table.Todo.AllColumns).
		ORDER_BY(table.Todo.ID.DESC()).
		LIMIT(int64(pageable.Size)).
		OFFSET(int64(pageable.Page) * int64(pageable.Size))

	countStmt := table.Todo.SELECT(postgres.COUNT(postgres.STAR)).FROM(table.Todo)

	var count struct { // 결과를 담을 구조체
		Count int64 `alias:"count"`
	}

	err2 := countStmt.Query(r.db, &count)

	if err2 != nil {
		errUtil.Wrap(err2)
	}

	var todos []model.Todo
	err := stmt.QueryContext(ctx, r.db, &todos)

	if err != nil {
		errUtil.Wrap(err)
	}

	result := dto.PageResult[model.Todo]{Content: todos, Size: pageable.Size, Total: int(count.Count)}

	return result
}

func (r *todoRepository) UpdateTodoStatus(ctx context.Context, id int32) (model.Todo, error) {

	stmt := table.Todo.UPDATE(table.Todo.Completed).
		SET(postgres.Raw("NOT completed")).
		WHERE(table.Todo.ID.EQ(postgres.Int32(id))).
		RETURNING(table.Todo.AllColumns) // 모든 컬럼 반환

	var todo model.Todo
	err := stmt.QueryContext(ctx, r.db, &todo)

	if err != nil {
		return todo, errUtil.Wrap(err)
	}
	// 업데이트된 행이 없는 경우
	if todo.ID == 0 {
		return todo, errUtil.Wrap(errors.New("todo id not found"))
	}
	return todo, nil
}
