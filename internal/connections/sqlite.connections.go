package connections

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnetion() *sql.DB {
	dbName := "dev.db"
	_, err := os.Create(dbName)
	if err != nil {
		if err != os.ErrExist {
			panic(err)
		}
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}
	return db
}
