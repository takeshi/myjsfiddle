package model

import (
	"github.com/coopernurse/gorp"
	"myjsfiddle/server/db"
)

func init() {
	dbMap := db.DbMap()
	Migrate(dbMap)
}

func Migrate(dbmap *gorp.DbMap) {
	dbmap.AddTable(Theme{}).SetKeys(true, "Id")
	dbmap.AddTable(Contents{}).SetKeys(true, "Id")
}
