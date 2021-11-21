package connections

import (
	"github.com/go-rel/rel"
	"github.com/go-rel/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteConection(dbName string) rel.Repository {
	adapter, err := sqlite3.Open(dbName)
	if err != nil {
		panic(err)
	}

	return rel.New(adapter)
}
