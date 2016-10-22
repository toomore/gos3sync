package syncdb

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	// go-sqlite3
	_ "github.com/mattn/go-sqlite3"
)

const indexDBFilename = "index.db"

// ConnDB connect to db
func ConnDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/%s", path, indexDBFilename))
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

// InsertIndex save data into index
func InsertIndex(db *sql.DB, path string, hex string) {
	tx, _ := db.Begin()
	tx.Prepare(`insert into idx(path, hex) values(?, ?)`)
	tx.Exec(path, hex)
	tx.Commit()
}

// Init init all
func Init(path string) {
	var err error
	path, err = filepath.Abs(path)
	if err != nil {
		log.Fatalln(err)
	}
	initTable(ConnDB(path))
	log.Printf("Create index db at '%s/%s'\n", path, indexDBFilename)
}
