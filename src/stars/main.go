package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	. "repository/employees"
	. "repository/geography"
	. "repository/regulators"
)

const (
	StarsDatabaseName = "stars"
)

func main() {
	m := martini.Classic()
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(martini.Static("site"))
	m.Use(sessions.Sessions("my_session", store))
	m.Use(render.Renderer(render.Options{Extensions: []string{".tmpl", ".html"}, Delims: render.Delims{"{{{", "}}}"}}))
	RegisterInjectors(m)
	RegisterRoutes(m)
	m.Run()
}

func RegisterInjectors(m *martini.ClassicMartini) {
	employeeRepository := NewEmployeeRepository(StarsDatabaseName)
	m.MapTo(employeeRepository, (*IEmployeeRepository)(nil))
	geographyRepository := NewGeographyRepository(StarsDatabaseName)
	m.MapTo(geographyRepository, (*IGeographyRepository)(nil))
	regulatorsRepository := NewRegulatorsRepository(StarsDatabaseName)
	m.MapTo(regulatorsRepository, (*IRegulatorsRepository)(nil))

}

type OperationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

func GotError(err error, rd render.Render, message ...string) bool {
	if err != nil {
		msg := err.Error()
		if len(message) > 0 {
			msg = fmt.Sprintf(message[0], err.Error())
		}
		rd.JSON(http.StatusOK, OperationResponse{
			StatusError, msg,
		})
		return true
	}
	return false
}
