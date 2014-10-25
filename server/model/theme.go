package model

import "myjsfiddle/server/db"

type Theme struct {
	Id       int64     `form:"id" db:"ThemeId"`
	Title    string    `form:"title"`
	Contents *Contents `db:"-"`
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

func (theme *Theme) Update() error {
	dbMap := db.DbMap()
	_, err := dbMap.Update(theme)
	if err != nil {
		return err
	}
	return theme.createContents()
}

func (theme *Theme) createContents() error {
	contents := theme.Contents
	if contents != nil {
		contents.ThemeId = theme.Id
		err := contents.Create()
		if err != nil {
			return err
		}
	}
	return nil
}

func (theme *Theme) Create() error {
	dbMap := db.DbMap()
	err := dbMap.Insert(theme)
	if err != nil {
		return err
	}
	return theme.createContents()
}

func GetTheme(themeId int64) (*Theme, error) {
	dbMap := db.DbMap()
	// theme := Theme{}
	obj, err := dbMap.Get(Theme{}, themeId)
	if obj == nil {
		return nil, err
	}
	return obj.(*Theme), err
}
