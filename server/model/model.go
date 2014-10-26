package model

import "myjsfiddle/server/db"

// type ModelImpl struct {
// 	Id int64
// }

type Model interface {
	GetId() int64
}

func CreateOrUpdate(model Model) error {
	dbMap := db.DbMap()
	if model.GetId() == 0 {
		return dbMap.Insert(model)
	} else {
		_, err := dbMap.Update(model)
		return err
	}

}

func Get(model Model, id int64) (interface{}, error) {
	dbMap := db.DbMap()
	return dbMap.Get(model, id)
}
