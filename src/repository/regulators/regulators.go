package regulators

import (
	. "domain"
	"errors"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"labix.org/v2/mgo/bson"
	"math"
	"storage"
	"time"
)

const (
	AssaysCollectionName      = "assays"
	AssayVisitsCollectionName = "assayvisits"
)

type IRegulatorsRepository interface {
	// Assays
	CreateMedicalAssay(id, name string, period int, active bool) (IMedicalAssay, error)
	AllMedicalAssays() ([]IMedicalAssay, error)
	FindMedicalAssay(assayId string) (IMedicalAssay, error)
	// Visits
	AllEmployeeAssaysLastVisits(login string) ([]IMedicalAssayVisit, error)
	UpdateEmployeeAssaysLastVisits(login string, visits []IMedicalAssayVisit) error
	// Statistics
	AllHotAssayVisits(date time.Time) ([]IMedicalAssayVisit, error)
}

//implementation

type regulatorsRepository struct {
	DatabaseName string
}

func NewRegulatorsRepository(databaseName string) IRegulatorsRepository {
	return &regulatorsRepository{
		DatabaseName: databaseName,
	}
}

func (repository *regulatorsRepository) AllMedicalAssays() ([]IMedicalAssay, error) {
	var err error
	db, err := storage.New(repository.DatabaseName)
	if err == nil {
		var list []*medicalAssay
		err = db.Collection(AssaysCollectionName).Find(nil).All(&list)
		if err == nil {
			models := make([]IMedicalAssay, len(list))
			for i, _ := range list {
				models[i] = IMedicalAssay(list[i])
			}
			return models, nil
		}
	}
	return nil, err
}

func (repository *regulatorsRepository) FindMedicalAssay(assayId string) (IMedicalAssay, error) {
	var err error
	db, err := storage.New(repository.DatabaseName)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var assay medicalAssay
	err = db.Collection(AssaysCollectionName).Find(bson.M{"_id": assayId}).One(&assay)
	if err != nil {
		return nil, err
	}
	return &assay, nil
}

func (repository *regulatorsRepository) CreateMedicalAssay(id, name string, period int, active bool) (IMedicalAssay, error) {
	var db storage.IDatabase
	var err error
	db, err = storage.New(repository.DatabaseName)
	if err == nil {
		defer db.Close()
		//find a duplicate
		var cnt int
		cnt, err = db.Collection(AssaysCollectionName).Find(bson.M{"_id": id}).Count()
		if err == nil {
			if cnt == 0 {
				c := NewMedicalAssay(id, name, period, active)
				err = db.Collection(AssaysCollectionName).Insert(c)
				if err == nil {
					return c, nil
				}
			} else {
				err = errors.New(fmt.Sprintf("The assay with id '%s' already exists", id))
			}
		}
	}
	return nil, err
}

const (
	VisitStatusAssayInactive = "inactive"
	VisitStatusNeverVisited  = "never"
	VisitStatusOk            = "ok"
	VisitStatusNotify        = "notify"
	VisitStatusWarning       = "warning"
	VisitStatusDangerous     = "dangerous"
	VisitStatusCritical      = "critical"
)

//
func (repository *regulatorsRepository) AllEmployeeAssaysLastVisits(login string) ([]IMedicalAssayVisit, error) {
	// for each assay
	assays, err := repository.AllMedicalAssays()
	if err != nil {
		return nil, err
	}
	db, err := storage.New(repository.DatabaseName)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var visits []*medicalAssayVisit
	err = db.Collection(AssayVisitsCollectionName).Find(nil).All(&visits)
	if err != nil {
		return nil, err
	}
	list := make([]IMedicalAssayVisit, 0, len(assays))
	for _, assay := range assays {
		//check whether a visit information is already exists
		e, found, err := From(visits).FirstBy(func(b T) (bool, error) {
			v := b.(IMedicalAssayVisit)
			return v.AssayId() == assay.Id() && v.EmployeeLogin() == login, nil
		})
		if err != nil {
			return nil, err
		}
		var visit IMedicalAssayVisit
		if !found {
			visit = NewMedicalAssayVisit(assay.Id(), login, time.Time{}, assay.Period())

			visit.Status(MedicalAssayVisitStatusNew)
		} else {
			visit = e.(IMedicalAssayVisit)
			visit.AssayPeriod(assay.Period())
			//Determine status for the visit
			if assay.Active() {
				if visit.Visited().IsZero() {
					if visit.Active() {
						visit.Status(VisitStatusNeverVisited)
						visit.Message("Осмотр не проводился")
					}
				} else if visit.Active() {
					status, msg := CalculateDaysDifference(visit.Visited(), assay.Period())
					visit.Status(status)
					visit.Message(msg)
				}
			} else {
				visit.Status(VisitStatusAssayInactive)
				visit.Message("Отменен")
			}
		}
		list = append(list, visit)
	}
	return list, nil
}

func CalculateDaysDifference(visited time.Time, period int) (status, message string) {
	planned := visited.AddDate(0, period, 0)
	elapsed := planned.Sub(time.Now())
	days := int(elapsed.Hours() / 24)
	message = fmt.Sprintf("%v дней", math.Abs(float64(days)))
	switch {
	case days > 14:
		status = VisitStatusOk
	case days > 7:
		status = VisitStatusNotify
	case days > 0:
		status = VisitStatusWarning
	case days > -3:
		status = VisitStatusDangerous
	default:
		status = VisitStatusCritical
	}
	return
}

// Updates employee medical assays visit information
func (repository *regulatorsRepository) UpdateEmployeeAssaysLastVisits(login string, visits []IMedicalAssayVisit) error {
	// 1.Get all the employee visits from db
	initialVisits, err := repository.AllEmployeeAssaysLastVisits(login)
	if err != nil {
		fmt.Println("E001 - Could not get all employee visits", err.Error())
		return err
	}
	// for each changed visit
	for _, v := range visits {
		//find the correcponding visit from db
		vs, found, err := From(initialVisits).FirstBy(func(b T) (bool, error) {
			return b.(IMedicalAssayVisit).AssayId() == v.AssayId(), nil
		})
		if err != nil {
			fmt.Println("E002 - First by linq error", err.Error())
			return err
		}
		db, err := storage.New(repository.DatabaseName) //^ err
		if err != nil {
			return err
		}
		defer db.Close()
		var visit IMedicalAssayVisit
		//create new visits if ther's no exists before
		if !found {
			visit = NewMedicalAssayVisit(v.AssayId(), v.EmployeeLogin(), v.Visited(), v.AssayPeriod())
			visit.Active(v.Active())
			visit.Status(MedicalAssayVisitStatusNew)
			fmt.Printf("After creating a visit: %v \n", visit)

		} else {
			visit = vs.(IMedicalAssayVisit)
			visit.Visited(v.Visited())
			visit.Active(v.Active())
			fmt.Printf("After casting a visit: %v \n", visit)
		}

		if visit.Status() == MedicalAssayVisitStatusNew {
			err := db.Collection(AssayVisitsCollectionName).Insert(visit)
			if err != nil {
				fmt.Println("E003 - Insert error", err.Error())
				return err
			}
		} else {
			//update visit
			visit.Active(v.Active())
			visit.Visited(v.Visited())
			fmt.Printf("After changing a visit: %v \n", visit)
			err = db.Collection(AssayVisitsCollectionName).Update(
				bson.M{"login": v.EmployeeLogin(), "assay": v.AssayId()},
				visit)
			if err != nil {
				fmt.Println("E004 - Update had not found", err.Error())
				return err
			}

		}

	}
	return nil

}

func (repository *regulatorsRepository) AllHotAssayVisits(date time.Time) ([]IMedicalAssayVisit, error) {
	return nil, nil
}
