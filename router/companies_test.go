package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/alcmoraes/go-data-integration-challenge/database"
	"github.com/alcmoraes/go-data-integration-challenge/types"
	"github.com/alcmoraes/go-data-integration-challenge/utils"
	"github.com/stretchr/testify/assert"

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

func TestCompanyAPI(t *testing.T) {

	R := Load()

	t.Run("Will fail on getting companies without sending parameters", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/companies/", nil)
		assert.NoError(t, err)
		R.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Returns 404 on company not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", fmt.Sprintf("/companies/?name=%s&zip=%d", "facetruck", 12345), nil)
		assert.NoError(t, err)
		R.ServeHTTP(w, req)
		log.Print(w.Body)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Success when persisting a CSV in database", func(t *testing.T) {

		values := map[string]io.Reader{
			"file":    utils.MustOpen("../test/mock/q1_catalog.csv"),
			"persist": strings.NewReader("true"),
		}

		w := httptest.NewRecorder()

		b, m, err := utils.GenerateMultipartHeadersFromLocal(values)

		req, err := http.NewRequest("POST", "/companies/upload", b)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", m.FormDataContentType())
		R.ServeHTTP(w, req)
		assert.NoError(t, err)
		log.Print(w.Body)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Returns a company when it finds one", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", fmt.Sprintf("/companies/?name=%s&zip=%d", "tola", 78229), nil)
		assert.NoError(t, err)
		R.ServeHTTP(w, req)
		log.Print(w.Body)
		assert.Equal(t, http.StatusOK, w.Code)
		company := types.Company{}
		err = json.Unmarshal(w.Body.Bytes(), &company)
		assert.NoError(t, err)
		assert.Equal(t, "tola sales group", company.Name)
		assert.Empty(t, company.Website)
	})

	t.Run("Success when merging a CSV in database", func(t *testing.T) {

		values := map[string]io.Reader{
			"file":    utils.MustOpen("../test/mock/q2_clientData.csv"),
			"persist": strings.NewReader("false"),
		}

		w := httptest.NewRecorder()

		b, m, err := utils.GenerateMultipartHeadersFromLocal(values)

		req, err := http.NewRequest("POST", "/companies/upload", b)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", m.FormDataContentType())
		R.ServeHTTP(w, req)
		assert.NoError(t, err)
		log.Print(w.Body)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Returns a merged company when it finds one", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", fmt.Sprintf("/companies/?name=%s&zip=%d", "tola", 78229), nil)
		assert.NoError(t, err)
		R.ServeHTTP(w, req)
		log.Print(w.Body)
		assert.Equal(t, http.StatusOK, w.Code)
		company := types.Company{}
		err = json.Unmarshal(w.Body.Bytes(), &company)
		assert.NoError(t, err)
		assert.Equal(t, "http://repsources.com", company.Website)
	})

}
