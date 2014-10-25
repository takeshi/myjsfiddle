package web

import "github.com/go-martini/martini"
import "github.com/martini-contrib/render"
import "github.com/martini-contrib/binding"
import "myjsfiddle/server/model"
import "myjsfiddle/server/db"

import "strconv"

// import "log"

func ThemeController(m *martini.ClassicMartini) {

	m.Get("/app/theme/:themeId/:contentsId", func(params martini.Params, r render.Render) {
		themeId, _ := strconv.ParseInt(params["themeId"], 10, 64)
		contentsId, _ := strconv.ParseInt(params["contentsId"], 10, 64)
		theme, _ := model.GetTheme(themeId)
		if theme == nil {
			r.JSON(404, "Not Found")
			return
		}
		contents, _ := model.GetContents(contentsId)
		if contents != nil {
			theme.Contents = contents
		}

		r.JSON(200, theme)
	})

	m.Post("/app/theme", binding.Bind(model.Theme{}), func(theme model.Theme, r render.Render) {

		db.Transaction(func() error {
			var err error
			if theme.Id == 0 {
				err = theme.Create()
			} else {
				err = theme.Update()
			}
			if err != nil {
				r.JSON(500, err)
				return err
			}
			r.JSON(200, theme)
			return nil
		})
	})

}
