package employees

import (
	. "domain"
	"errors"
	"fmt"
	"labix.org/v2/mgo/bson"
	"storage"
)

//Employees Repository
type IEmployeeRepository interface {
	// Finds an employee by login name
	FindByLogin(login string) (IEmployee, error)
	// Create a new employee
	CreateNew(lastName, firstName, middleName, login, password string, createdBy IEmployee) (IEmployee, error)
	// Updates the employee
	SaveEmployee(employee IEmployee) error
	// Get Employees list
	AllEmployees() ([]IEmployee, error)
	// Authenticates an employee by login and password specified
	AuthenticateEmployee(login, password string) (IEmployee, error)
	// Get All Employee Roles
	AllEmployeeRoles() []IEmployeeRole
	//
	RoleIdsTitle(roleIds []string) (string, error)
	//
	FindEmployeeRoleById(roleId string) (IEmployeeRole, error)
}

// -----------------------------------------------------------------------------

const (
	EmployeesCollectionName = "employees"
)

type employeeRepository struct {
	DatabaseName  string
	EmployeeRoles []IEmployeeRole
}

func NewEmployeeRepository(databaseName string) IEmployeeRepository {
	return &employeeRepository{
		DatabaseName: databaseName,
	}
}

// Find an employee by login
func (repository *employeeRepository) FindByLogin(login string) (IEmployee, error) {
	db, err := storage.New(repository.DatabaseName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	emp := &employee{}
	err = db.Collection(EmployeesCollectionName).Find(bson.M{"_id": login}).One(emp)
	if err != nil {
		return nil, errors.New("employee " + login + " not found, reason - " + err.Error())
	}
	return emp, nil
}

// Create a new employee
func (repository *employeeRepository) CreateNew(lastName, firstName, middleName, login, password string, createdBy IEmployee) (IEmployee, error) {
	emp := NewEmployee(lastName, firstName, middleName, login, createdBy)
	emp.Password(password)
	db, err := storage.New(repository.DatabaseName)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Collection(EmployeesCollectionName).Insert(emp)
	if err != nil {
		return nil, err
	}
	return emp, nil
}

// Save employee instance
func (repository *employeeRepository) SaveEmployee(employee IEmployee) error {
	db, err := storage.New(repository.DatabaseName)
	if err != nil {
		return err
	}
	err = db.Collection(EmployeesCollectionName).Update(bson.M{"_id": employee.Login()}, employee)
	return err
}

// Get all employees registered in the db
func (repository *employeeRepository) AllEmployees() ([]IEmployee, error) {
	db, err := storage.New(repository.DatabaseName)
	if err != nil {
		return nil, err
	}
	var list []employee
	err = db.Collection(EmployeesCollectionName).Find(nil).All(&list)
	if err != nil {
		return nil, err
	}
	models := make([]IEmployee, len(list))
	for i, _ := range list {
		models[i] = IEmployee(&list[i])
	}
	return models, nil
}

//Authenticates employer by login and password specified
func (repository *employeeRepository) AuthenticateEmployee(login, password string) (IEmployee, error) {
	db, err := storage.New(repository.DatabaseName)
	if err != nil {
		return nil, err
	}
	var emp employee
	err = db.Collection(EmployeesCollectionName).Find(bson.M{"_id": login}).One(&emp)
	if err != nil {
		return nil, err
	}
	if !emp.Active() {
		return nil, errors.New("User is disabled")
	}
	if emp.Password() != password {
		return nil, errors.New("User found but not authenticated")
	}
	return &emp, nil
}

// Get All Employee Roles
func (repository *employeeRepository) AllEmployeeRoles() []IEmployeeRole {
	if len(repository.EmployeeRoles) == 0 {
		repository.EmployeeRoles = append(repository.EmployeeRoles,
			NewEmployeeRole("ADM", "Администратор"),
			NewEmployeeRole("DSP", "Диспетчер"),
			NewEmployeeRole("COK", "Повар"),
			NewEmployeeRole("DRV", "Водитель"),
			NewEmployeeRole("PRS", "Персонал"),
			NewEmployeeRole("TME", "Временщик"))
	}
	return repository.EmployeeRoles
}

func (repository *employeeRepository) RoleIdsTitle(roleIds []string) (string, error) {
	var s string
	for _, id := range roleIds {
		role, err := repository.FindEmployeeRoleById(id)
		if err != nil {
			return "", err
		}
		if len(s) > 0 {
			s += ", "
		}
		s += role.Name()
	}
	return s, nil
}

func (repository *employeeRepository) FindEmployeeRoleById(roleId string) (IEmployeeRole, error) {
	for i, v := range repository.AllEmployeeRoles() {
		if v.Id() == roleId {
			return repository.EmployeeRoles[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("No employee role with id '%s' found", roleId))
}
