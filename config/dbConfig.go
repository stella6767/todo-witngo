package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func LoadDB() *sql.DB {

	fmt.Println(GlobalConfig.Datasource.DbType)
	//dsn := "postgresql://localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open(GlobalConfig.Datasource.DbType, GlobalConfig.Datasource.Url)
	db.SetMaxOpenConns(10)

	if err != nil {
		panic(err)
	}

	return db
}
