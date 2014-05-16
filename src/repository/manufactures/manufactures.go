package manufactures

import (
	. "domain"
	//	"errors"
	//	"labix.org/v2/mgo/bson"
	"storage"
)

type IManufacturesRepository interface {
	AllManufactures() ([]IManufacture, error)
	CreateNew(name string, address IAddress) (IManufacture, error)
}

const (
	ManufacturesCollectionName = "manufactures"
)

type manufacturesRepository struct {
	DatabaseName string
}

func NewManufacturesRepository(databaseName string) IManufacturesRepository {
	return &manufacturesRepository{
		DatabaseName: databaseName,
	}
}

func (repository *manufacturesRepository) CreateNew(name string, address IAddress) (IManufacture, error) {
	man := NewManufacture(name, address)
	db, err := storage.New(repository.DatabaseName)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Collection(ManufacturesCollectionName).Insert(man)
	if err != nil {
		return nil, err
	}
	return man, nil
}

func (repository *manufacturesRepository) AllManufactures() ([]IManufacture, error) {
	db, err := storage.New(repository.DatabaseName)
	if err != nil {
		return nil, err
	}
	var list []*manufacture
	err = db.Collection(ManufacturesCollectionName).Find(nil).All(&list)
	models := make([]IManufacture, len(list))
	for i, v := range list {
		models[i] = IManufacture(v)
	}
	return models, nil
}
