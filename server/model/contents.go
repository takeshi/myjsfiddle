package model

type Contents struct {
	Id        int64 `db:"ContentsId"`
	ThemeId   int64
	Seq       int64
	Html      string
	Js        string
	Css       string
	Temp      bool
	Directive string
	Sass      string `db:"-"`
}

func (model *Contents) GetId() int64 {
	return model.Id
}
