package main

import (
	"github.com/codegangsta/martini"
	bind "github.com/martini-contrib/binding"
)

func RegisterRoutes(m *martini.ClassicMartini) {

	// Common web
	m.Get("/", authorize, RootWeb) // Root page
	m.NotFound(authorize, RootWeb) // Navigate when no path found

	// Authentication
	m.Get("/login", Entrance)                               // Login page
	m.Get("/logout", Exit)                                  // Logout
	m.Post("/login", bind.Form(LoginFormModel{}), TryLogin) // Login procedure

	//Dashboard
	m.Get("/dashboard/page", authorize, DashboardPage) // Get dashboard page

	//Employee
	m.Get("/employees/page", authorize, EmployeesPage)                                    // Gets all employees page
	m.Get("/employees/list", authorize, EmployeesList)                                    // Gets list of all employees
	m.Get("/employee/data/@", authorize, NewEmployeeData)                                 //Gets employee data
	m.Get("/employee/data/:login", authorize, EmployeeData)                               //Gets employee data
	m.Get("/employee/page", authorize, EmployeePage)                                      //Gets employee new/edit page
	m.Post("/employee/create", authorize, bind.Json(EmployeeDataModel{}), EmployeeCreate) //Gets employee new/edit page
	m.Post("/employee/update", authorize, bind.Json(EmployeeDataModel{}), EmployeeUpdate) //Gets employee new/edit page
	m.Get("/employee/profile", authorize, EmployeeProfilePage)                            //Gets employee profile page

	//Assays
	m.Get("/assays/page", authorize, AssaysPage)                                // Gets assays page
	m.Get("/assays/list/:login", authorize, AssaysList)                         // Gets assays list
	m.Post("/assays/update", authorize, bind.Json(VisitsModel{}), AssaysUpdate) // Update assays list

}
