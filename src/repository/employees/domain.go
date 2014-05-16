package employees

import (
	. "domain"
	"fmt"
	c "repository/common"
	"time"
)

type employeeRole struct {
	c.Unique
	c.Named
}

func NewEmployeeRole(id, name string) IEmployeeRole {
	role := &employeeRole{}
	role.IdΩ = id
	role.NameΩ = name
	return role
}

func NewEmployee(lastName, firstName, middleName, login string, createdBy IEmployee) IEmployee {
	now := time.Now()
	var employee = &employee{RolesΩ: make([]string, 0, 10)}
	employee.LastName(lastName)
	employee.FirstName(firstName)
	employee.MiddleName(middleName)
	employee.Login(login)
	employee.Active(false)
	employee.CreatedΩ = now
	employee.CreatedByΩ = createdBy
	employee.JobStarted(now)
	return employee
}

type employee struct {
	c.Activable `bson:"activable"`
	c.Trackable `bson:"trackable"`
	LoginΩ      string    `bson:"_id"`
	PasswordΩ   string    `bson:"password"`
	RolesΩ      []string  `bson:"roles"`
	FirstNameΩ  string    `bson:"firstName"`
	LastNameΩ   string    `bson:"lastName"`
	MiddleNameΩ string    `bson:"middleName"`
	GenderIdΩ   string    `bson:"gender"`
	BirthDateΩ  time.Time `bson:"birthdate"`
	JobStartedΩ time.Time `bson:"jobStarted"`
	MarkΩ       int       `bson:"mark"`
	AddressΩ    string    `bson:"address"`
	CreditΩ     bool      `bson:"credit"`
	EmailΩ      string    `bson:"email"`
	PhoneΩ      string    `bson:"phone"`
	CountryIdΩ  string    `bson:"country"`
}

// IUnique implementation
func (this employee) Id() string {
	return this.LoginΩ
}

// INameable implementation
func (this employee) Name(Ω ...string) string {
	if len(Ω) > 0 {
		panic("Employer Name property isn't assignable")
	}
	return fmt.Sprintf("%s %s", this.LastName(), this.FirstName())
}

// FirstName
func (this *employee) FirstName(Ω ...string) string {
	if len(Ω) > 0 {
		this.FirstNameΩ = Ω[0]
	}
	return this.FirstNameΩ
}

// LastName
func (this *employee) LastName(Ω ...string) string {
	if len(Ω) > 0 {
		this.LastNameΩ = Ω[0]
	}
	return this.LastNameΩ
}

// MiddleName
func (this *employee) MiddleName(Ω ...string) string {
	if len(Ω) > 0 {
		this.MiddleNameΩ = Ω[0]
	}
	return this.MiddleNameΩ
}

// Login
func (this *employee) Login(Ω ...string) string {
	if len(Ω) > 0 {
		this.LoginΩ = Ω[0]
	}
	return this.LoginΩ
}

//Password
func (this *employee) Password(Ω ...string) string {
	if len(Ω) > 0 {
		this.PasswordΩ = Ω[0]
	}
	return this.PasswordΩ
}

//RoleIds
func (this *employee) RoleIds() []string {
	roles := make([]string, len(this.RolesΩ))
	copy(roles, this.RolesΩ)
	return roles
}

const RoleAbsentIndex = -1

func roleIdIndex(roles []string, roleId string) int {
	for i, v := range roles {
		if v == roleId {
			return i
		}
	}
	return RoleAbsentIndex
}

//AssignRoleIds
func (this *employee) AssignRoleIds(roleIds ...string) {
	for i, _ := range roleIds {
		if roleIdIndex(this.RolesΩ, roleIds[i]) == RoleAbsentIndex {
			this.RolesΩ = append(this.RolesΩ, roleIds[i])
		}
	}
}

//IsInRole
func (this *employee) IsInRole(roleId string) bool {
	return roleIdIndex(this.RolesΩ, roleId) != RoleAbsentIndex
}

//Gender
func (this *employee) GenderId(Ω ...string) string {
	if len(Ω) > 0 {
		this.GenderIdΩ = Ω[0]
	}
	return this.GenderIdΩ
}

//BirthDate
func (this *employee) BirthDate(Ω ...time.Time) time.Time {
	if len(Ω) > 0 {
		this.BirthDateΩ = Ω[0]
	}
	return this.BirthDateΩ
}

//JobStarted
func (this *employee) JobStarted(Ω ...time.Time) time.Time {
	if len(Ω) > 0 {
		this.JobStartedΩ = Ω[0]
	}
	return this.JobStartedΩ
}

//Mark
func (this *employee) Mark(Ω ...int) int {
	if len(Ω) > 0 {
		this.MarkΩ = Ω[0]
	}
	return this.MarkΩ
}

//Address
func (this *employee) Address(Ω ...string) string {
	if len(Ω) > 0 {
		this.AddressΩ = Ω[0]
	}
	return this.AddressΩ
}

//Credit
func (this *employee) Credit(Ω ...bool) bool {
	if len(Ω) > 0 {
		this.CreditΩ = Ω[0]
	}
	return this.CreditΩ
}

//Email
func (this *employee) Email(Ω ...string) string {
	if len(Ω) > 0 {
		this.EmailΩ = Ω[0]
	}
	return this.EmailΩ
}

//Phone
func (this *employee) Phone(Ω ...string) string {
	if len(Ω) > 0 {
		this.PhoneΩ = Ω[0]
	}
	return this.PhoneΩ
}

//Country
func (this *employee) CountryId(Ω ...string) string {
	if len(Ω) > 0 {
		this.CountryIdΩ = Ω[0]
	}
	return this.CountryIdΩ
}

func (this *employee) ReAssignRoleIds(Ω ...string) {
	this.RolesΩ = make([]string, len(Ω))
	copy(this.RolesΩ, Ω)
}
