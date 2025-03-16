package test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"todo-app/config"
	"todo-app/internal/dto"
	"todo-app/internal/repository"
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

	result := config.InitAppDependency()
	fmt.Print(result)
}

func TestViper(t *testing.T) {

	config.LoadConfig()

}

func TestRepository(t *testing.T) {

	db, err := sql.Open("postgres", "postgresql://localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	todoRepository := repository.NewTodoRepository(db)

	pageable := dto.Pageable{Page: 0, Size: 10}
	ctx := context.Background() // 기본 컨텍스트 생성
	todos := todoRepository.GetTodosByPage(ctx, pageable)

	fmt.Println("???")
	fmt.Println(todos.Total)

	for i := 0; i < len(todos.Content); i++ {
		fmt.Println(todos.Content[i].Title)
	}

}
