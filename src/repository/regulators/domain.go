package regulators

import (
	. "domain"
	. "repository/common"
	"time"
)

type medicalAssay struct {
	IdΩ       string `bson:"_id"`
	Named     `bson:"named"`
	Activable `bson:"activable"`
	PeriodΩ   int `bson:"period"`
}

func (this medicalAssay) Id() string {
	return this.IdΩ
}

func (this *medicalAssay) Period(Ω ...int) int {
	if len(Ω) > 0 {
		this.PeriodΩ = Ω[0]
	}
	return this.PeriodΩ
}

func NewMedicalAssay(id, name string, period int, active bool) IMedicalAssay {
	a := &medicalAssay{}
	a.IdΩ = id
	a.Name(name)
	a.Period(period)
	a.Active(active)
	return a
}

type medicalAssayVisit struct {
	Activable      `bson:"activable"`
	AssayIdΩ       string    `bson:"assay"`
	EmployeeLoginΩ string    `bson:"login"`
	VisitedΩ       time.Time `bson:"visited"`
	StatusΩ        string    `bson:"-"`
	MessageΩ       string    `bson:"-"`
	AssayPeriodΩ   int       `bson:"-"`
}

const (
	MedicalAssayVisitStatusNew = "NewVisit"
)

func NewMedicalAssayVisit(assayId, employeeLogin string, visited time.Time, assayPeriod int) IMedicalAssayVisit {
	visit := &medicalAssayVisit{
		AssayIdΩ:       assayId,
		EmployeeLoginΩ: employeeLogin,
		VisitedΩ:       visited,
		StatusΩ:        AssayVisitOk,
		AssayPeriodΩ:   assayPeriod,
	}
	visit.Active(false)
	return visit
}

func (this *medicalAssayVisit) AssayId() string {
	return this.AssayIdΩ
}
func (this *medicalAssayVisit) AssayPeriod(Ω ...int) int {
	if len(Ω) > 0 {
		this.AssayPeriodΩ = Ω[0]
	}
	return this.AssayPeriodΩ
}
func (this *medicalAssayVisit) EmployeeLogin() string {
	return this.EmployeeLoginΩ
}

func (this *medicalAssayVisit) Visited(Ω ...time.Time) time.Time {
	if len(Ω) > 0 {
		this.VisitedΩ = Ω[0]
	}
	return this.VisitedΩ
}

func (this *medicalAssayVisit) Status(Ω ...string) string {
	if len(Ω) > 0 {
		this.StatusΩ = Ω[0]
	}
	return this.StatusΩ
}

func (this *medicalAssayVisit) Message(Ω ...string) string {
	if len(Ω) > 0 {
		this.MessageΩ = Ω[0]
	}
	return this.MessageΩ
}
