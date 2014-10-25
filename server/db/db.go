package db

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"log"
	"os"

	// "myjsfiddle/server/model"

	_ "github.com/mattn/go-sqlite3"
)

var dbmap *gorp.DbMap
var filePath = "myjsfiddle.db"

func Transaction(f func() error) error {
	dbmap := DbMap()
	tx, err := dbmap.Begin()
	if err != nil {
		log.Print(err, "Transaction Begin Error")
		return err
	}
	err = f()
	if err != nil {
		log.Print(err, "Transaction RError")
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func init() {
	dbmap = CreateDbMap(filePath)
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

	return dbmap
}

// func migrate(dbmap *gorp.DbMap) {
// 	dbmap.AddTable(model.Theme{}).SetKeys(true, "Id")
// 	dbmap.AddTable(model.Contents{}).SetKeys(true, "Id")
// }

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
