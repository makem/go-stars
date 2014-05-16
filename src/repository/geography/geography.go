package geography

import (
	. "domain"
	"errors"
	"fmt"
	"labix.org/v2/mgo/bson"
	"storage"
)

const (
	CountryCollectionName string = "countries"
)

type IGeographyRepository interface {
	NewCountry(code, name string) (ICountry, error)
	AllCountries() ([]ICountry, error)
}

//implementation
type geographyRepository struct {
	DatabaseName string
}

//
func NewGeographyRepository(databaseName string) IGeographyRepository {
	return &geographyRepository{
		DatabaseName: databaseName,
	}
}

func (repository *geographyRepository) NewCountry(code, name string) (ICountry, error) {
	var db storage.IDatabase
	var err error
	db, err = storage.New(repository.DatabaseName)
	if err == nil {
		defer db.Close()
		//find a duplicate
		var cnt int
		cnt, err = db.Collection(CountryCollectionName).Find(bson.M{"_id": code}).Count()
		if err == nil {
			if cnt == 0 {
				c := newCountry(code, name)
				err = db.Collection(CountryCollectionName).Insert(c)
				if err == nil {
					return c, nil
				}
			} else {
				err = errors.New(fmt.Sprintf("The country with code '%s' already exists", code))
			}
		}
	}
	return nil, err

}

//
func (repository *geographyRepository) AllCountries() ([]ICountry, error) {
	var err error
	db, err := storage.New(repository.DatabaseName)
	if err == nil {
		var list []*country
		err = db.Collection(CountryCollectionName).Find(nil).All(&list)
		if err == nil {
			models := make([]ICountry, len(list))
			for i, _ := range list {
				models[i] = ICountry(list[i])
			}
			return models, nil
		}
	}
	return nil, err
}
