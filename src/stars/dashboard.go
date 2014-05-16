package main

import (
	"github.com/martini-contrib/render"
	"net/http"
)

func DashboardPage(rd render.Render) {
	rd.HTML(http.StatusOK, "pages/dashboard", nil)
}
