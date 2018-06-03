package importer

import (
	"bytes"
	"encoding/csv"
	"io"
	"mime/multipart"

	"github.com/alcmoraes/go-data-integration-challenge/database"
	"github.com/alcmoraes/go-data-integration-challenge/types"
	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
)

// Worker for the CSV file import data
func Worker(f multipart.File, h *multipart.FileHeader, persist bool, done chan bool) error {

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.LazyQuotes = true
		r.Comma = ';'
		return r
	})

	companies := []*types.Company{}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, f); err != nil {
		log.Error(err)
		return err
	}

	err := gocsv.UnmarshalBytes(buf.Bytes(), &companies)
	if err != nil {
		log.Error(err)
		return err
	}

	for _, c := range companies {
		company := *c
		go database.AddCompanyIntoDatabase(company, persist)

	}
	done <- true
	return nil
}
