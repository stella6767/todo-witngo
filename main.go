package main

import (
	"database/sql"
	"fmt"
	"github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
	"todo-app/.gen/postgres/public/model"
	"todo-app/.gen/postgres/public/table"
)

func main() {

	//cfg, err := config.LoadConfig()
	//if err != nil {
	//	log.Fatal("Failed to load config:", err)
	//}
	//
	//// DB 연결
	//pool, err := pgxpool.New(context.Background(), cfg.DBURL)
	//if err != nil {
	//	log.Fatal("Failed to connect to DB:", err)
	//}
	//defer pool.Close()
	//
	//// Repository 초기화
	//queries := db.New(pool)
	//todoRepo := repository.NewTodoRepository(queries)
	//
	//router.Router(todoRepo, cfg)

	// Connect to database
	//var connectString = fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable", "localhost", "5432", "public")
	//fmt.Print(connectString)
	dsn := "postgresql://localhost:5432/postgres?sslmode=disable"
	db, err :=
		sql.Open("postgres", dsn)

	panicOnError(err)
	defer db.Close()

	// Write query
	stmt := postgres.SELECT(
		table.Todo.AllColumns,
	).FROM(table.Todo)

	var dest []model.Todo
	stmt.Query(db, &dest)
	panicOnError(err)

	dest = append(dest, model.Todo{Title: "test"})
	fmt.Println(dest)

	var peopleSlice []Person
	peopleSlice = append(peopleSlice, Person{Name: "John"})
	fmt.Println(peopleSlice)

}

type Person struct {
	Name string
	Age  int
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
