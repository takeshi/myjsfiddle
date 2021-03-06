package model

import "testing"
import "myjsfiddle/server/db"

import "myjsfiddle/server/model"
import "log"

func TestContents(t *testing.T) {
	dbmap := db.CreateDbMap("Test.db")
	model.Migrate(dbmap)
	defer dbmap.DropTablesIfExists()
	theme := model.Theme{
		Title: "Sample",
	}
	model.CreateOrUpdate(&theme)

}

type GetIder interface {
	GetId() int64
}

type A struct {
	Id int64
}

type B struct {
	A
	Id int64
}

func (a *A) GetId() int64 {
	return a.Id
}

func Test2(t *testing.T) {
	b := B{}
	b.Id = 100
	log.Println(b, b.GetId())
	var a GetIder
	a = &b
	log.Println(a, a.GetId())

	t.Fail()
}
