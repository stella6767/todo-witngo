package test

import (
	"fmt"
	"testing"
	"todo-app/config"
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

	loadConfig, err := config.LoadConfig()
	if err != nil {
		fmt.Errorf("Fatal error config file: %s \n", err)
	}

	fmt.Print(loadConfig)

}
