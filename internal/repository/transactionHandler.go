package repository

import (
	"context"
	"database/sql"
	errUtil "todo-app/internal/errUtils"
)

type TransactionHandler struct {
	db *sql.DB
}

func NewTxHandler(db *sql.DB) *TransactionHandler {
	return &TransactionHandler{db: db}
}

func (t *TransactionHandler) Execute(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return errUtil.Wrap(err)
	}
	//`tx.Commit()`을 호출하면 트랜잭션이 종료되기 때문에, 이후에 `tx.Rollback()`을 호출해도 이미 커밋된 트랜잭션에는 영향을 주지 않음.
	//그래서 `defer`로 롤백을 등록해놓아도, 커밋이 성공하면 롤백은 무시
	defer tx.Rollback()

	// 트랜잭션 컨텍스트 생성
	txCtx := context.WithValue(ctx, "tx", tx)

	// 3. 비즈니스 로직 실행
	err2 := fn(txCtx)

	// 4. 에러 발생 시 롤백
	if err2 != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errUtil.Wrap(rbErr)
		}
	}

	// 5. 커밋 처리
	if err := tx.Commit(); err != nil {
		return errUtil.Wrap(err)
	}
	return nil
}
