package model

type Theme struct {
	Id       int64     `form:"id" db:"ThemeId"`
	Title    string    `form:"title"`
	Contents *Contents `db:"-"`
}

func (model *Theme) GetId() int64 {
	return model.Id
}

func (theme *Theme) CreateContents(html, js, css string) *Contents {
	contents := Contents{
		Html:    html,
		Css:     css,
		Js:      js,
		ThemeId: theme.Id,
	}
	return &contents
}
