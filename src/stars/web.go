package main

import (
	. "domain"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	//	"github.com/martini-contrib/binding"
	"log"
	"net/http"
	. "repository/employees"
)

type IndexModel struct {
	Employee IEmployee
}

func RootWeb(rd render.Render, s sessions.Session, lg *log.Logger, employeeRepository IEmployeeRepository) {
	login := s.Get("login").(string)
	employee, err := employeeRepository.FindByLogin(login)
	if err != nil {
		rd.Error(500)
	} else {
		rd.HTML(http.StatusOK, "index", &IndexModel{
			Employee: employee,
		})
	}
}
