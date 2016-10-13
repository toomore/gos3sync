package syncdb

import (
	"database/sql"
	"fmt"
	"log"

	// go-sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// ConnDB connect to db
func ConnDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/index.db", path))
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func initTable(db *sql.DB) {
	_, err := db.Exec(`create table idx (path TEXT NOT NULL PRIMARY KEY, hex TEXT)`)
	if err != nil {
		log.Fatalln(err)
	}
}

// Init init all
func Init(path string) {
	initTable(ConnDB(path))
}
