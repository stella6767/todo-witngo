package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// or 싱글톤 패턴을 적용해서 repository 생성자에서 주입하는게 아니라, 전역에서 갖고오게 하는 함수를 만드는 방법도
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
