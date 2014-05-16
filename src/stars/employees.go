package main

import (
	. "domain"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"net/http"
	. "repository/employees"
	. "repository/geography"
	"time"
)

func EmployeesPage(rd render.Render) {
	rd.HTML(http.StatusOK, "pages/employees", nil)
}

type EmployeePageData struct {
	ContactInfo  string
	PersonalInfo string
	PositionInfo string
	FirstName    string
	MiddleName   string
	LastName     string
	Man          string
	Woman        string
	BirthDate    string
	RegisterDate string
	Email        string
	Phone        string
	Country      string
	Address      string
	Save         string
	Cancel       string
	Actual       string
	Credit       string
	Login        string
}

func EmployeePage(rd render.Render) {
	rd.HTML(http.StatusOK, "pages/employee", &EmployeePageData{
		ContactInfo:  "Контактная информация",
		PersonalInfo: "Персональные данные",
		PositionInfo: "Регистрационная информация",
		FirstName:    "Имя",
		MiddleName:   "Отчество",
		LastName:     "Фамилия",
		Man:          "Муж",
		Woman:        "Жен",
		BirthDate:    "Дата рождения",
		RegisterDate: "Дата приема на работу",
		Email:        "Емейл",
		Phone:        "Телефон",
		Country:      "Страна",
		Address:      "Адрес проживания",
		Save:         "Сохранить",
		Cancel:       "Отмена",
		Actual:       "Действующий",
		Credit:       "Кредит",
		Login:        "Логин",
	})
}

func EmployeeProfilePage(rd render.Render) {
	rd.HTML(http.StatusOK, "pages/profile", nil)
}

//
type EmployeeModel struct {
	Name       string `json:"name"`
	Login      string `json:"login"`
	Active     bool   `json:"active"`
	Roles      string `json:"roles"`
	Registered string `json:"registered"`
	Gender     string `json:"gender"`
}

//
func EmployeesList(rd render.Render, repository IEmployeeRepository) {
	list, err := repository.AllEmployees()
	if err != nil {
		rd.Error(http.StatusInternalServerError)
	}
	res := make([]*EmployeeModel, len(list))
	for i, v := range list {
		title, err := repository.RoleIdsTitle(v.RoleIds())
		if err != nil {
			title = err.Error()
		}
		res[i] = &EmployeeModel{
			Name:       v.Name(),
			Login:      v.Login(),
			Active:     v.Active(),
			Registered: v.JobStarted().Format(DateShortFormat),
			Roles:      title,
			Gender:     v.GenderId(),
		}
	}
	rd.JSON(http.StatusOK, res)
}

type RoleModel struct {
	Name   string `json:"name"`
	Id     string `json:"id"`
	Active bool   `json:"active"`
}

type EmployeeDataAnswerModel struct {
	Model     EmployeeDataModel `json:"model"`
	Countries []ICountry        `json:"countries"`
}

type EmployeeDataModel struct {
	Id         string       `json:"id"`
	Login      string       `json:"login"`
	LastName   string       `json:"lastName"`
	FirstName  string       `json:"firstName"`
	MiddleName string       `json:"middleName"`
	Active     bool         `json:"active"`
	Roles      []*RoleModel `json:"roles"`
	Gender     string       `json:"gender"`
	BirthDate  string       `json:"birthDate"`
	JobStarted string       `json:"jobStarted"`
	Mark       int          `json:"mark"`
	Address    string       `json:"address"`
	Credit     bool         `json:"credit"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	Country    string       `json:"country"`
}

// Translates
func computeRolesModel(roleIds []string, repository IEmployeeRepository) ([]*RoleModel, error) {
	roles := repository.AllEmployeeRoles()
	model := make([]*RoleModel, 0, len(roles))
	for _, v := range repository.AllEmployeeRoles() {
		any, _ := From(roleIds).AnyWith(func(s T) (bool, error) {
			return s.(string) == v.Id(), nil
		})
		model = append(model, &RoleModel{
			Id:     v.Id(),
			Name:   v.Name(),
			Active: any,
		})
	}
	return model, nil
}

func NewEmployeeData(rd render.Render, repository IEmployeeRepository, geo IGeographyRepository) {
	roleIds := make([]string, 0)
	roles, err := computeRolesModel(roleIds, repository)
	if err != nil {
		rd.Error(http.StatusInternalServerError)
	}
	var birthDate string
	var jobStarted string
	model := EmployeeDataModel{
		Roles:      roles,
		BirthDate:  birthDate,
		JobStarted: jobStarted,
		Gender:     MaleGenderId,
	}
	countries, err := geo.AllCountries()
	if err != nil {
		rd.Error(http.StatusInternalServerError)
	}
	dataModel := EmployeeDataAnswerModel{
		Model:     model,
		Countries: countries,
	}
	rd.JSON(http.StatusOK, dataModel)
}

//
func EmployeeData(rd render.Render, repository IEmployeeRepository, geo IGeographyRepository, params martini.Params) {
	login := params["login"]
	if len(login) == 0 {
		rd.Error(http.StatusBadRequest)
	}
	emp, err := repository.FindByLogin(login)
	if err != nil {
		rd.Error(http.StatusInternalServerError)
	}
	roles, err := computeRolesModel(emp.RoleIds(), repository)
	if err != nil {
		rd.Error(http.StatusInternalServerError)
	}
	var birthDate string
	var jobStarted string
	if !emp.BirthDate().IsZero() {
		birthDate = emp.BirthDate().Format(DateShortFormat)
	}
	if !emp.JobStarted().IsZero() {
		jobStarted = emp.JobStarted().Format(DateShortFormat)
	}
	model := EmployeeDataModel{
		Id:         login,
		Login:      login,
		LastName:   emp.LastName(),
		FirstName:  emp.FirstName(),
		MiddleName: emp.MiddleName(),
		Active:     emp.Active(),
		Roles:      roles,
		Mark:       emp.Mark(),
		BirthDate:  birthDate,
		JobStarted: jobStarted,
		Address:    emp.Address(),
		Credit:     emp.Credit(),
		Email:      emp.Email(),
		Phone:      emp.Phone(),
		Gender:     emp.GenderId(),
		Country:    emp.CountryId(),
	}
	countries, err := geo.AllCountries()
	if err != nil {
		rd.Error(http.StatusInternalServerError)
	}
	dataModel := EmployeeDataAnswerModel{
		Model:     model,
		Countries: countries,
	}
	rd.JSON(http.StatusOK, dataModel)
}

func parseRolesModel(roles []*RoleModel) []string {
	list := make([]string, 0, len(roles))
	for i, v := range roles {
		if v.Active {
			list = append(list, roles[i].Id)
		}
	}
	return list
}

func UpsertEmployee(e IEmployee, m *EmployeeDataModel) error {
	e.Active(m.Active)
	e.Address(m.Address)
	bd, err := time.Parse(DateShortFormat, m.BirthDate)
	if err != nil {
		return err
	}
	e.BirthDate(bd)
	rd, err := time.Parse(DateShortFormat, m.JobStarted)
	if err != nil {
		return err
	}
	e.JobStarted(rd)
	e.CountryId(m.Country)
	e.Credit(m.Credit)
	e.Email(m.Email)
	e.FirstName(m.FirstName)
	e.GenderId(m.Gender)
	e.LastName(m.LastName)
	e.Mark(m.Mark)
	e.MiddleName(m.MiddleName)
	e.Phone(m.Phone)
	r := parseRolesModel(m.Roles)
	e.ReAssignRoleIds(r...)
	return nil
}

//
func EmployeeCreate(rd render.Render, repository IEmployeeRepository, m EmployeeDataModel) {
	emp, err := repository.CreateNew(m.LastName, m.FirstName, m.MiddleName, m.Login, "12345", nil)
	if err != nil {
		rd.JSON(http.StatusOK, OperationResponse{
			Status:  StatusError,
			Message: err.Error(),
		})
		return
	}
	UpsertEmployee(emp, &m)
	err = repository.SaveEmployee(emp)
	if err != nil {
		rd.JSON(http.StatusOK, OperationResponse{
			Status:  StatusError,
			Message: err.Error(),
		})
		return
	}
	rd.JSON(http.StatusOK, OperationResponse{
		Status: StatusSuccess,
	})
}

//
func EmployeeUpdate(rd render.Render, repository IEmployeeRepository, m EmployeeDataModel) {
	emp, err := repository.FindByLogin(m.Id)
	if err != nil {
		rd.JSON(http.StatusOK, OperationResponse{
			Status:  StatusError,
			Message: fmt.Sprintf("Can't find a user with login %s because of error: %s", m.Id, err.Error()),
		})
		return
	}
	UpsertEmployee(emp, &m)
	err = repository.SaveEmployee(emp)
	if err != nil {
		rd.JSON(http.StatusOK, OperationResponse{
			Status:  StatusError,
			Message: err.Error(),
		})
		return
	}
	rd.JSON(http.StatusOK, OperationResponse{
		Status: StatusSuccess,
	})
}
