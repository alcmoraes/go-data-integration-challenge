package database

import (
	"errors"
	"fmt"

	"github.com/alcmoraes/go-data-integration-challenge/config"
	"github.com/alcmoraes/go-data-integration-challenge/types"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

type DatabaseOptions struct {
	Testing bool
}

var Options DatabaseOptions

func GetSession() (*mgo.Session, error) {
	session, err := mgo.Dial(config.Config.DBAddress)
	return session, err
}

func GetDatabase() (*mgo.Session, *mgo.Database, error) {
	session, err := GetSession()
	if err != nil {
		return &mgo.Session{}, &mgo.Database{}, fmt.Errorf("Not possible to connect to %s database: %v", config.Config.Database, err)
	}
	db := session.DB(config.Config.Database)
	return session, db, nil
}

func GetCollection(collection string) (*mgo.Session, *mgo.Collection, error) {
	session, db, err := GetDatabase()
	if err != nil {
		return &mgo.Session{}, &mgo.Collection{}, fmt.Errorf("Not possible to connect to %s collection: %v", collection, err)
	}
	c := db.C(collection)
	return session, c, nil
}

func AddCompanyIntoDatabase(c types.Company, persist bool) error {
	session, collection, err := GetCollection("companies")
	defer session.Close()
	if err != nil {
		log.Error(err)
		return err
	}

	if len(c.Zip) != 5 {
		return errors.New("Zip must be 5 chars length")
	}

	companies := []*types.Company{}
	err = collection.Find(bson.M{"name": bson.M{"$regex": "^" + c.Name + "\\s?.*", "$options": "i"}, "zip": c.Zip}).All(&companies)
	if err != nil {
		log.Error(err)
		return err
	}

	if len(companies) > 1 {
		return errors.New("Multiple companies match the given data.")
	}

	if len(companies) == 1 && !persist {
		err = collection.UpdateId(companies[0].ID, c)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	if len(companies) == 0 && persist {
		err = collection.Insert(c)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

func init() {
	Options = DatabaseOptions{
		Testing: false,
	}
}
