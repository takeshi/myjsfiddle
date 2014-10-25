package model

import "testing"
import "myjsfiddle/server/db"
import "myjsfiddle/server/model"

func TestContents(t *testing.T) {
	dbmap := db.CreateDbMap("Test.db")
	defer dbmap.DropTablesIfExists()
	theme := model.Theme{
		Title: "Sample",
	}
	_ = dbmap.Insert(&theme)

	contents := theme.CreateContents("SampleHtml", "SampleJS", "SampleCSS")
	if theme.Id != contents.ThemeId {
		t.Errorf("expected %v actual %v", theme.Id, contents.ThemeId)
	}
}
