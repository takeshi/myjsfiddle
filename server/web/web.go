package web

import "github.com/go-martini/martini"

func Init(m *martini.ClassicMartini) {
	RunController(m)
	ThemeController(m)
}
