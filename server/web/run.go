package web

import "github.com/go-martini/martini"
import "github.com/martini-contrib/render"
import _ "github.com/martini-contrib/binding"
import "myjsfiddle/server/db"
import "myjsfiddle/server/model"
import "html/template"
import "log"
import "github.com/moovweb/gosass"

func RunController(m *martini.ClassicMartini) {

	m.Get("/fiddle/app/run/:themeId/:id", func(params martini.Params, r render.Render) {
		dbMap := db.DbMap()
		contents := model.Contents{}
		err := dbMap.SelectOne(&contents, "select * from Contents where ThemeId=? and ContentsId=?", params["themeId"], params["id"])
		if err != nil {
			log.Print(err, "select contents error")
			r.JSON(200, "No Contents")
			return
		}
		obj, _ := model.Get(&model.Theme{}, contents.ThemeId)
		theme := obj.(*model.Theme)
		context := gosass.Context{
			Options: gosass.Options{
				OutputStyle:    gosass.COMPRESSED_STYLE,
				SourceComments: false,
			},
			SourceString: contents.Css,
		}
		gosass.Compile(&context)
		r.HTML(200, "hello", map[string]interface{}{
			"Title":     template.HTML(theme.Title),
			"Css":       template.CSS(context.OutputString),
			"Js":        template.JS(contents.Js),
			"Html":      template.HTML(contents.Html),
			"Directive": template.JS(contents.Directive),
		})
	})
}
