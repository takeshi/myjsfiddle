package model

import (
	"github.com/coopernurse/gorp"
	"log"
	// "myjsfiddle/server/db"
)

// func init() {
// 	dbMap := db.DbMap()
// 	Migrate(dbMap)
// }

func Migrate(dbmap *gorp.DbMap) {
	dbmap.AddTable(Theme{}).SetKeys(true, "Id")
	dbmap.AddTable(Contents{}).SetKeys(true, "Id")
	err := dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}
}
