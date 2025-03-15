package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func loadDB() *sql.DB {
	dsn := "postgresql://localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}
