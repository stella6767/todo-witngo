package test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"testing"
	"todo-app/.gen/postgres/public/model"
	"todo-app/config"
	"todo-app/internal/dto"
	"todo-app/internal/errUtils"
	"todo-app/internal/repository"
	"todo-app/internal/util"
)

/** 작성규칙
파일명은 _test.go 로 끝난다.
함수명은 Test 로 시작한다.
매개변수는 *testing.T를 받는다.
실패지점에서 t.Fail()을 호출한다.

모든 테스트 실행 -> $ go test
*_test.go 대상으로 테스트 실행하며 상세한 수행내역( -v ) 확인 $ go test *_test.go -v\
특정 테스트 함수(TestFunctionName) 실행 $ go test -run TestFunctionName -v
*/

func TestInitApp(t *testing.T) { // 테스트

	result := config.InitAppDependency(nil)
	fmt.Print(result)
}

func TestViper(t *testing.T) {

	config.LoadConfig("")

}

func TestGetTodosByPage(t *testing.T) {

	_, todoRepository := createTestTodoRepository()

	pageable := dto.Pageable{Page: 0, Size: 10}
	ctx := context.Background() // 기본 컨텍스트 생성
	todos, _ := todoRepository.GetTodosByPage(ctx, pageable)

	fmt.Println(todos.Total)

	for i := 0; i < len(todos.Content); i++ {
		todo := todos.Content[i]
		//todo.ID

		if *todo.Completed == true {

		}

		fmt.Println(todo.Title)
	}

}

func loadTestDb() *sql.DB {

	//db, err := sql.Open("postgres", "postgresql://localhost:5432/postgres?sslmode=disable")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		"localhost",
		"5432",
		"postgres",
		"1234",
		"postgres",
		"disable",
		"public",
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	return db
}

func TestUpdateTodoStatus(t *testing.T) {
	//ctx, todoRepository := createTestTodoRepository()
	//
	//_, err = errorWrapTest(err, todoRepository, ctx)
	//
	//if err != nil {
	//	fmt.Printf("%+v\n", err)
	//	log.WithError(err).Error("", err)
	//}

}

func TestCreateTodo(t *testing.T) {
	ctx, todoRepository := createTestTodoRepository()

	todo, err := todoRepository.CreateTodo(ctx, "task???")

	if err != nil {
		log.WithError(err).Error("", err)
	}

	fmt.Println(todo)
}

func createTestTodoRepository() (context.Context, repository.TodoRepository) {
	ctx := context.Background() // 기본 컨텍스트 생성
	db := loadTestDb()

	log.SetFormatter(&util.PrettyFormatter{
		TextFormatter: log.TextFormatter{
			ForceColors:  true,
			PadLevelText: true,
		},
	})

	todoRepository := repository.NewTodoRepository(db)
	return ctx, todoRepository
}

func errorWrapTest(err error, todoRepository repository.TodoRepository, ctx context.Context) (model.Todo, error) {
	result, err := todoRepository.UpdateTodoStatus(ctx, 1)
	if err != nil {
		return result, errUtil.Wrap(err)
	}
	return result, nil
}

func TestError(t *testing.T) {
	err := c()
	if err != nil {

		fmt.Printf("%+v\n", err)
	}
}

func a() error {
	return errors.New("first errUtils")
}
func b() error {
	return errors.Wrap(a(), "second errUtils")
}
func c() error {
	return errors.Wrap(b(), "third errUtils")
}

func TestErrorUtil(t *testing.T) {

	log.SetFormatter(&util.PrettyFormatter{
		TextFormatter: log.TextFormatter{
			ForceColors:  true,
			PadLevelText: true,
		},
	})

	if err := foo(); err != nil {
		err = errUtil.Wrap(err) // 추가 message가 필요 없을 때
		//fmt.Printf("%+v\n", err)
		//log.Errorf("%+v", err)
		//log.Errorf("%+v\n", err)

		log.Error(err)

		//log.WithError(err).Warn("db close error")

	}
}

func foo() error {
	if err := bar(); err != nil { // 하위 함수에서 errUtils 발생
		return errUtil.WrapWithMessage(err, "foo message") // 추가 message가 필요할 때
	}
	return nil
}

func bar() error {
	return errors.New("bar Error!")
}
