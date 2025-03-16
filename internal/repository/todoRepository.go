package repository

import (
	"context"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"log"
	"todo-app/.gen/postgres/public/model"
	"todo-app/.gen/postgres/public/table"
	"todo-app/internal/dto"
)

// 공개되는 인터페이스
type TodoRepository interface {
	CreateTodo(ctx context.Context, title string) (*model.Todo, error)
	GetTodos(ctx context.Context) ([]model.Todo, error)
	GetTodosByPage(ctx context.Context, pageable dto.Pageable) dto.PageResult[model.Todo]
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
		log.Fatal(err2.Error())
	}
	var todos []model.Todo
	err := stmt.QueryContext(ctx, r.db, &todos)

	if err != nil {
		log.Fatal(err.Error())
	}
	
	return dto.PageResult[model.Todo]{Content: todos, Size: pageable.Size, Total: int(count.Count)}
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
