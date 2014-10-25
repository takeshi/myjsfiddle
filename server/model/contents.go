package model

import "myjsfiddle/server/db"

type Contents struct {
	Id      int64 `db:"ContentsId"`
	ThemeId int64
	Seq     int64
	Html    string
	Js      string
	Css     string
	Temp    bool
}

func (contens *Contents) Create() error {
	dbmap := db.DbMap()
	return dbmap.Insert(contens)
}

func GetContents(contentsId int64) (*Contents, error) {
	dbMap := db.DbMap()
	// theme := Theme{}
	obj, err := dbMap.Get(Contents{}, contentsId)
	if obj == nil {
		return nil, err
	}
	return obj.(*Contents), err
}
