package connections

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteConection(dbName string) *sql.DB {
	conn, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}
	return conn
}
