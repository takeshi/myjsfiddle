package db

import "testing"
import "./"

func TestSqrt(t *testing.T) {
	dbmap := db.CreateDbMap("Test.db")
	defer dbmap.DropTablesIfExists()

}
