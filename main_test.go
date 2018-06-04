package main_test

import (
	"os"
	"testing"

	"github.com/alcmoraes/go-data-integration-challenge/database"
	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {

	session, db, err := database.GetDatabase()
	defer session.Close()

	if err != nil {
		log.Error(err)
	}

	err = db.DropDatabase()
	if err != nil {
		log.Error(err)
	}

	os.Exit(m.Run())
}
