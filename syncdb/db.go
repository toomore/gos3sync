package syncdb

import (
	"database/sql"
	"log"

	// go-sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// ConnDB connect to db
func ConnDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./index.db")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func initTable() {
	var db = ConnDB()
	_, err := db.Exec(`create table idx (path TEXT NOT NULL PRIMARY KEY, hex TEXT)`)
	if err != nil {
		log.Fatalln(err)
	}
}

// Init init all
func Init() {
	initTable()
}
