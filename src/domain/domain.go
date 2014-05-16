package domain

import (
	"time"
)

type IUnique interface {
	Id() string
}

type INamed interface {
	Name(Ω ...string) string
}

type IActivable interface {
	Active(Ω ...bool) bool
}

type ITrackable interface {
	Created() time.Time
	CreatedBy() IEmployee
}

type IEntity interface {
	IUnique
	INamed
	IActivable
	ITrackable
}

type IEmployeeRole interface {
	IUnique
	INamed
}

type IGender interface {
	IUnique
	INamed
}

type IEmployee interface {
	IEntity
	FirstName(Ω ...string) string
	LastName(Ω ...string) string
	MiddleName(Ω ...string) string
	GenderId(Ω ...string) string
	BirthDate(Ω ...time.Time) time.Time
	RoleIds() []string
	AssignRoleIds(Ω ...string)
	ReAssignRoleIds(Ω ...string)
	IsInRole(roleId string) bool
	JobStarted(Ω ...time.Time) time.Time
	Mark(Ω ...int) int
	Address(Ω ...string) string
	Credit(Ω ...bool) bool
	Email(Ω ...string) string
	Phone(Ω ...string) string
	CountryId(Ω ...string) string
	Login(Ω ...string) string
	Password(Ω ...string) string
}

type ICountry interface {
	IUnique
	INamed
}

type IAddress interface {
	Country(Ω ...ICountry) ICountry
	City(city ...string) string
	District(district ...string) string
	Street(street ...string) string
	House(house ...string) string
	Building(building ...string) string
	Appartment(appartment ...string) string
}

type IManufacture interface {
	IUnique
	INamed
	IActivable
	Address(address ...IAddress) IAddress
	Opened() time.Time
}

type IMedicalAssay interface {
	IUnique
	INamed
	IActivable
	Period(Ω ...int) int
}

const (
	AssayVisitOk       = "ok"
	AssayVisitWarning  = "warning"
	AssayVisitOutdated = "outdated"
)

type IMedicalAssayVisit interface {
	IActivable
	AssayId() string
	AssayPeriod(Ω ...int) int
	EmployeeLogin() string
	Visited(Ω ...time.Time) time.Time
	Status(Ω ...string) string
	Message(Ω ...string) string
}
