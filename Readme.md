
# Demo

[go-todo](https://gotodo.freeapp.me/)

# go path preSetup

```azure
~/GolandProjects/todo-app main !1 ❯ sudo chown -R $(whoami):staff /Users/{username}/go      
~/GolandProjects/todo-app main !1 ❯ chmod -R 755 /Users/{username}/go
```

# test stack

* GIN
* templ
* postgres
* air
* viper
* go-jet
* htmx
* tailwind



# Http Framework

Gin과 Echo 크게 두가지가 대중적으로 쓰고 있는듯. 여기서는 GIN을 선택
[embed로 스태틱 리소스 넣기](https://medium.com/bgpworks/golang-1-16%EC%97%90-%EC%83%88%EB%A1%9C-%EC%B6%94%EA%B0%80%EB%90%9C-%EA%B8%B0%EB%8A%A5-embed%EB%A1%9C-%EC%8A%A4%ED%83%9C%ED%8B%B1-%EB%A6%AC%EC%86%8C%EC%8A%A4-%EB%84%A3%EA%B8%B0-1675c4564f5e)
[A Guide to Embedding Static Files in Go using go:embed](https://www.iamyadav.com/blogs/a-guide-to-embedding-static-files-in-go)


# UI Rendering

## templ

최대한 html과 가까운 서버사이드 템플릿 엔진을 찾아보는데 templ 이 가장 가까운 것 같다. 좀 아쉬운 부분이 많다. lsp plugin을 깔아서 쓰는데, 자동완성 및, auto import, 오류 구문 표시 같은 게 거의 없는 수준이고, 
문법 오류 표시가 없어서 generate 하고 메시지를 보고 다시 찾아가서 고치는 경우가 많다. 찾아봐도 대안이 없는데.. Go 진영에서는 이게 최선인가?  

### templ cli

![img.png](img.png)

```azure
export PATH="$PATH:/usr/local/go/bin/bin"
```

공식 문서에서는 air와 연동해서 Hot reloading을 하는 방법을 소개해주고 있는데, 되긴 되는데 도중에 templ 문법을 틀려도 잡아주질 않는다. 뭔가 불완전 

[Go + HTMX + Templ + Tailwind: 프로젝트 설정 완료 및 핫 리로딩](https://medium.com/ostinato-rigore/go-htmx-templ-tailwind-complete-project-setup-hot-reloading-2ca1ba6c28be)


# DB Handling

## sqlc

[Getting started with PostgreSQL](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html)

[sqlc overrides](https://docs.sqlc.dev/en/stable/howto/overrides.html)

[여러분은 pgx pgtype 보일러플레이트를 어떻게 다루고 계신가요?](https://www.reddit.com/r/golang/comments/1h5q7ng/how_are_you_guys_dealing_with_pgx_pgtype/)

```azure
# example
overrides:
   - db_type: "pg_catalog.int4"
     nullable: true
     go_type:
       type:  "int"
```

실제로 써보니 쿼리빌더라기보다, 원시 sql 작성하면 그걸 기반으로 repository 계층 코드를 generate 해주는 코드생성도구에 가깝다.
JVM에서의 JOOQ 같은 걸 기대했는데, 생각보다 더 원시적이다. 런타임에 동적인 코드 생성이 불가능? 한 것 같다. 더 알아보아야 겠지만,
특히 짜증나는 건, postgres 쓰면 pgtype을 import 해줘야 되는데, 이걸 go native type으로 변환시키는 작업을 일일히 해야 된다는 점이다.
위의 overrides 설정을 적용하면 일부 해결이 되는 것 같은데, 또 안 되는 것도 있고 하.. 짜증이 난다. 일단 폐기처분.


## go-jet

[go-jet](https://github.com/go-jet/jet?tab=readme-ov-file#features)

```azure
$ go get -u github.com/go-jet/jet/v2
go install github.com/go-jet/jet/v2/cmd/jet@latest
jet -dsn='postgresql://localhost:5432/postgres?sslmode=disable' -schema=public -path=./.gen
```

![img_1.png](img_1.png)

JOOQ 랑 가장 비슷한 형식의 라이브러리인 듯 하다. 이걸로 정하자.

### 트랜잭션 handling

스프링처럼 선언적으로 트랜잭션을 구현 또는 비슷하게 하는 법이 있을까 찾아보다가 영상을 하나 보게 되었는데
[Transaction Management and Repository Pattern | Ilia Sergunin | Conf42 Golang 2023](https://www.youtube.com/watch?v=aRsea6FFAyA&ab_channel=Conf42)
고차 함수를 통해서 트랜잭션 AOP를 구현하고 있다. 좋은 생각인 것 같아서, transaction manger inerface를 만들고 
바로 따라 해볼려고 했으나 문제가..
```azure
Interface method must have no type parameters
```
Go에서는 interface 함수는 제네릭 타입을 가질 수 없다. 그래서 인터페이스 선언부에 Generice 를 선언하려 했으나..

```azure
type TransactionHandler[T any] interface {
	Execute(ctx context.Context, fn func(ctx context.Context) (T, error)) (T, error)
}
```
인터페이스가 인스턴스화될 때 인터페이스의 제네릭 유형을 지정해야 되었다. 따라서 인터페이스 생성 시점부터 제네릭의 유형이
바인딩되고 변경불가다. 이러면 유연하게 여기저기서 쓸 수가 없다. 나는 하나의 인스턴스로 돌려 쓰고 싶은데, 매번 타입이 달라질 
때마다 새로 만들 수는 없는 노릇 아닌가.. 구조체도 마찬가지다. 
관련 재밌는 논의링크 하나 남긴다.

[제안: 사양: 메서드에서 유형 매개변수 허용 #49085](https://github.com/golang/go/issues/49085)






# DI Container in Go?

[wire](https://github.com/google/wire)

[uber/dig](https://github.com/uber-go/dig)

[inject](https://github.com/facebookarchive/inject)

찾아보니.. 크게 위의 3가지 라이브러리가 주로 쓰이고 있는 것 같다. Go 진영에서는 DI 라이브러리를 그닥 좋아하지는 않는 것 같다. 가급적 수동 DI를 유지하기로 결정  

# 환경변수 관리법

[godotenv](https://github.com/joho/godotenv)

[viper](https://github.com/spf13/viper)

크게 2가지 정도 많이 쓰고 있는 것 같은데, 다양한 포맷과 형식을 지원하는 viper 사용


# 프로젝트 구조

[Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ko.md)

[대규모 Gin-GORM 웹 서비스 구성: 효과적인 폴더 구조 가이드](https://fenixara.com/organizing-a-large-scale-gin-gorm-web-service-a-guide-to-effective-folder-structure/)

[go-gin-boilerplate](https://github.com/vsouza/go-gin-boilerplate)

사실 여러 개 문서를 뒤져봐도 딱히 그럴싸한 구조를 선정하지 못하겠다. 그냥 나름대로 Spring mvc와 비슷하게 구성했다.

# GRPC

# 컨벤션

일단 참고용

[뱅크샐러드 Go 코딩 컨벤션](https://blog.banksalad.com/tech/go-best-practice-in-banksalad/)


# Error Handling

go는 기본적으로 stack trace 정보를 제공해주지 않는다. 처음 봤을 시 조금 충격적이었는데, 아래 글을 많이 참고
[Golang, 그대들은 어떻게 할 것인가 - 4. error 핸들링](https://d2.naver.com/helloworld/6507662)
[panic-and-recover-more](https://go101.org/article/panic-and-recover-more.html)

# build && execute
```azure
uname -m && uname -s # 배포할 서버의 os 및 아키텍처 확인 

GOOS=linux GOARCH=arm64 go build -o todoapp # 정적 바이너리 파일 빌드
    
scp -i ~/.ssh/{filename}.pem todoapp prod.yaml {user}@{ip}:~/cicd/go-todo    
    
 ./todoapp # 실행 
nohup env GO_PROFILE=prod ./todoapp > todoapp.log 2>&1 & 

```


# 후기

간결하고 빠른 언어임에는 분명하지만 프로트엔드 작업하기에는 좋지 않다. 유연하고 풍부한 표현가능한 API가 부족하고
(특히 타입캐스팅 측면) 생태계 측면에서도 부족한 게 많다. 그럴싸한 서버사이드 템플릿 엔진 중에 가장 성숙한 게 templ
같은데 수준이 너무 떨어진다. 그래도 나온지 10년이 넘은 언어인데 이 정도가 최선인가? 