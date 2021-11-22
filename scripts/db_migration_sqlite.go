package scripts

import (
	"context"
	"log"
	"os"

	"github.com/go-rel/rel"
	"github.com/manicar2093/web-monitor/db/dao"
	"github.com/manicar2093/web-monitor/db/entities"
)

type MigrationToSqlite struct {
	Conn rel.Repository
	// PageDaoFile must be *dao.PageDaoImpl type
	PageDaoFile dao.PageDao
	// PageDaoSqlite must be *dao.PageDaoSqlite type
	PageDaoSqlite dao.PageDao
	// PhraseDaoFile must be *dao.PhraseDaoImpl type
	PhraseDaoFile dao.PhraseDao
	// PhraseDaoSqlite must be *dao.PhraseDaoSqlite type
	PhraseDaoSqlite dao.PhraseDao
}

const (
	findTableCreatedSQL = `
		SELECT 
			name
		FROM sqlite_schema
		WHERE
			type = "table" AND
			name = ?
	`
)

var (
	logFile = func() *os.File {
		f, err := os.OpenFile("migration_to_sqlite.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644)

		if err != nil {
			log.Println("Error opening logger file")
			panic(err)
		}
		return f
	}
	logger log.Logger = *log.New(logFile(), "STATE", log.LstdFlags)
)

func DatabaseMigrationToSqlite(depend *MigrationToSqlite) {
	validateDependsTypes(depend)
	migratePages(depend)
	migratePhrases(depend)

}

func migratePages(depend *MigrationToSqlite) {
	if tableExist(entities.Page{}.Table(), depend) {
		return
	}
	_, _, err := depend.Conn.Exec(context.Background(), entities.PageCreationTableSQL)
	if err != nil {
		logger.Println("An unexpected error occurred creating Page Table")
		panic(err)
	}
	log.Println("Pages table created")
	pagesFromFile, err := depend.PageDaoFile.GetAllPages()
	if err != nil {
		logger.Println("Error getting all pages from DBV1")
		panic(err)
	}

	for _, page := range pagesFromFile {
		if err := depend.PageDaoSqlite.Save(page); err != nil {
			switch {
			case err == rel.ErrUniqueConstraint:
				continue
			default:
				logger.Printf("Error migrating '%s' page. Details: %v\n", page.Name, err)
			}
		}
	}
}

func migratePhrases(depend *MigrationToSqlite) {
	if tableExist(entities.Phrase{}.Table(), depend) {
		return
	}
	_, _, err := depend.Conn.Exec(context.Background(), entities.PhraseCreationTableSQL)
	if err != nil {
		logger.Println("An unexpected error occurred creating Phrase Table")
		panic(err)
	}
	log.Println("Phrase table created")
	phraseFromFile, err := depend.PhraseDaoFile.GetAllPhrases()
	if err != nil {
		logger.Println("Error getting all phrase from DBV1")
		panic(err)
	}

	for _, phrase := range phraseFromFile {
		if err := depend.PhraseDaoSqlite.Save(phrase); err != nil {
			switch {
			case err == rel.ErrUniqueConstraint:
				continue
			default:
				logger.Printf("Error migrating '%s' phrase. Details: %v\n", phrase.Phrase, err)
			}
		}
	}
}

func tableExist(tableName string, depend *MigrationToSqlite) bool {
	var tableData struct {
		Name string
	}
	sql := rel.SQL(findTableCreatedSQL, tableName)

	err := depend.Conn.Find(context.Background(), &tableData, sql)
	if err != nil {
		if err == rel.ErrNotFound {
			return false
		}
		logger.Printf("Error finding table created. Details: %v\n", err)
		panic(err)
	}
	return tableData.Name == tableName
}

func validateDependsTypes(depend *MigrationToSqlite) {
	if _, ok := depend.PageDaoFile.(*dao.PageDaoImpl); !ok {
		panic("PageDaoFile should be of type *dao.PageDaoImpl")
	}
	if _, ok := depend.PageDaoSqlite.(*dao.PageDaoSqlite); !ok {
		panic("PageDaoSqlite should be of type *dao.PageDaoSqlite")
	}
	if _, ok := depend.PhraseDaoFile.(*dao.PhraseDaoImpl); !ok {
		panic("PhraseDaoFile should be of type *dao.PhraseDaoImpl")
	}
	if _, ok := depend.PhraseDaoSqlite.(*dao.PhraseDaoSqlite); !ok {
		panic("PhraseDaoSqlite should be of type *dao.PhraseDaoSqlite")
	}

}
