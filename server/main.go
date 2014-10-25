package server

import "github.com/go-martini/martini"
import "github.com/martini-contrib/render"

// import "os"

import "myjsfiddle/server/web"
import "myjsfiddle/server/model"
import "myjsfiddle/server/db"

func Server() {
	// os.Setenv("Setenv", "8080")
	// port := os.Getenv("PORT")
	// println(port)
	dbMap := db.DbMap()
	model.Migrate(dbMap)
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(martini.Static("dist"))
	m.Use(martini.Static("./"))
	web.Init(m)
	m.Run()

}
