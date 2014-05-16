package storage

import (
	"errors"
	"fmt"
	"labix.org/v2/mgo"
)

const (
	ServerUri string = "mongodb://localhost"
)

type IDatabase interface {
	Collection(name string) *mgo.Collection
	Truncate() error
	Close()
}

type database struct {
	Session  *mgo.Session
	Database *mgo.Database
}

func TruncateDatabase(databaseName string) error {
	db, err := New(databaseName)
	if err != nil {
		return err
	}
	return db.Truncate()
}

func TruncateCollection(databaseName, collectionName string) error {
	db, err := New(databaseName)
	if err != nil {
		return err
	}
	return db.Collection(collectionName).DropCollection()
}

func New(name ...string) (IDatabase, error) {
	session, err := mgo.Dial(ServerUri)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Can't connect to mongo, go error %v\n", err))
	}
	databaseName := "Stars"
	if len(name) > 0 {
		databaseName = name[0]
	}
	session.SetSafe(&mgo.Safe{})
	db := session.DB(databaseName)
	if db == nil {
		return nil, errors.New("Cant get database " + databaseName)
	}
	return &database{
		Session:  session,
		Database: db,
	}, nil
}

func (db *database) Collection(name string) *mgo.Collection {
	c := db.Database.C(name)
	return c
}

func (db *database) Close() {
	db.Session.Close()
	db.Database = nil
	db.Session = nil
}

func (db *database) Truncate() error {
	return db.Database.DropDatabase()
}
