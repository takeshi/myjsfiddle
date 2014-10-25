package db

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var dbmap *gorp.DbMap
var filePath = "myjsfiddle.db"

func init() {
	dbmap = CreateDbMap(filePath)
}

func Transaction(f func() error) error {
	dbmap := DbMap()
	tx, err := dbmap.Begin()
	if err != nil {
		log.Print(err, "Transaction Begin Error")
		return err
	}
	err = f()
	if err != nil {
		log.Print(err, "Transaction Rollback")
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DbMap() *gorp.DbMap {
	return dbmap
}

func CreateDbMap(filePath string) *gorp.DbMap {

	db, err := sql.Open("sqlite3", filePath)
	checkErr(err, "sql.Open failed")

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	// migrate(dbmap)

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
