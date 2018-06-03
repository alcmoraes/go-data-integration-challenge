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

func GetSession() (*mgo.Session, error) {
	session, err := mgo.Dial(config.Config.DBAddress)
	return session, err
}

func GetCollection(collection string) (*mgo.Session, *mgo.Collection, error) {
	session, err := GetSession()
	if err != nil {
		return &mgo.Session{}, &mgo.Collection{}, fmt.Errorf("Not possible to connect to %s collection: %v", collection, err)
	}

	c := session.DB(config.Config.Database).C(collection)
	return session, c, nil
}

func AddCompanyIntoDatabase(c types.Company, persist bool) error {
	session, collection, err := GetCollection("companies")
	defer session.Close()
	if err != nil {
		log.Error(err)
		return err
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
