package web

import "github.com/go-martini/martini"
import "github.com/martini-contrib/render"
import "github.com/martini-contrib/binding"
import "myjsfiddle/server/model"
import "myjsfiddle/server/db"
import "net/http"

import "strconv"
import "log"
import "fmt"

func ThemeController(m *martini.ClassicMartini) {

	m.Get("/fiddle/app/theme/:themeId/:contentsId", func(params martini.Params, r render.Render) {
		themeId, _ := strconv.ParseInt(params["themeId"], 10, 64)
		contentsId, _ := strconv.ParseInt(params["contentsId"], 10, 64)
		m, err := model.Get(&model.Theme{}, themeId)
		if err != nil {
			log.Println(err)
			r.JSON(404, "Not Found")
			return
		}
		theme := m.(*model.Theme)
		m, err = model.Get(&model.Contents{}, contentsId)
		if err != nil {
			log.Println(err)
		} else {
			contents := m.(*model.Contents)
			theme.Contents = contents
		}
		log.Println(m)
		r.JSON(200, theme)
	})

	m.Post("/fiddle/app/theme", binding.Bind(model.Theme{}), func(theme model.Theme, r render.Render) {

		db.Transaction(func() error {
			var err error
			log.Println(theme.Id, theme.GetId())
			err = model.CreateOrUpdate(&theme)
			if err != nil {
				r.JSON(500, err)
				return err
			}
			if theme.Contents != nil {
				contents := theme.Contents
				contents.ThemeId = theme.Id
				contents.Id = 0
				err = model.CreateOrUpdate(contents)
				if err != nil {
					r.JSON(500, err)
					return err
				}
			}
			r.JSON(200, theme)
			return nil
		})
	})

	m.Post("/fiddle/app/theme/fork", func(req *http.Request, r render.Render) {

		db.Transaction(func() error {
			var err error
			theme := model.Theme{
				Title: req.FormValue("title"),
			}
			contents := model.Contents{
				Css:       req.FormValue("css"),
				Js:        req.FormValue("js"),
				Html:      req.FormValue("html"),
				Directive: req.FormValue("directive"),
			}
			err = model.CreateOrUpdate(&theme)
			if err != nil {
				r.JSON(500, err)
				return err
			}
			contents.ThemeId = theme.Id
			err = model.CreateOrUpdate(&contents)
			if err != nil {
				r.JSON(500, err)
				return err
			}
			r.Redirect(fmt.Sprintf("/fiddle/#/fiddle/%d/%d/%s", theme.Id, contents.Id, req.FormValue("mode")))
			return nil
		})
	})
}
