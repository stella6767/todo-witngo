package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func loadDB() *sql.DB {

	fmt.Println(globalConfig.Datasource.DbType)

	//dsn := "postgresql://localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open(globalConfig.Datasource.DbType, globalConfig.Datasource.Url)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}
