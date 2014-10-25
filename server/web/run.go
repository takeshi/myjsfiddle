package web

import "github.com/go-martini/martini"
import "github.com/martini-contrib/render"
import _ "github.com/martini-contrib/binding"
import "myjsfiddle/server/db"
import "myjsfiddle/server/model"
import "html/template"
import "log"

func RunController(m *martini.ClassicMartini) {

	m.Get("/app/run/:themeId/:id", func(params martini.Params, r render.Render) {
		dbMap := db.DbMap()
		contents := model.Contents{}
		err := dbMap.SelectOne(&contents, "select * from Contents where ThemeId=? and ContentsId=?", params["themeId"], params["id"])
		if err != nil {
			log.Print(err, "select contents error")
			r.JSON(200, "NotFound")
			return
		}
		r.HTML(200, "hello", map[string]interface{}{
			"Css":  template.CSS(contents.Css),
			"Js":   template.JS(contents.Js),
			"Html": template.HTML(contents.Html),
		})
	})
}
