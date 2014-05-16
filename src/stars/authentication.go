package main

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	. "repository/employees"
)

func authorize(session sessions.Session, rd render.Render) {
	if session.Get("login") == nil {
		rd.Redirect("/login")
	}
}

func Entrance(rd render.Render) {
	rd.HTML(http.StatusOK, "login", nil)
}

func Exit(s sessions.Session, rd render.Render) {
	s.Clear()
	rd.Redirect("/")
}

type LoginFormModel struct {
	Login    string `form:"login"`
	Password string `form:"password"`
}

//Try Login user into the system
func TryLogin(s sessions.Session,
	rd render.Render,
	employeeRepository IEmployeeRepository,
	model LoginFormModel) {
	_, err := employeeRepository.AuthenticateEmployee(model.Login, model.Password)
	if err == nil {
		s.Set("login", model.Login)
		rd.Redirect("/", http.StatusMovedPermanently)
	} else {
		rd.Redirect("/login")
	}
}
