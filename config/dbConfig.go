package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// or 싱글톤 패턴을 적용해서 repository 생성자에서 주입하는게 아니라, 전역에서 갖고오게 하는 함수를 만드는 방법도
func LoadDB() *sql.DB {

	datasource := GlobalConfig.Datasource

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		datasource.Host,
		datasource.Port,
		datasource.UserName,
		datasource.Password,
		datasource.Dbname,
		"disable",
		datasource.Schema,
	)

	db, err := sql.Open(datasource.DbType, connStr)
	db.SetMaxOpenConns(10)

	if err != nil {
		panic(err)
	}

	return db
}
