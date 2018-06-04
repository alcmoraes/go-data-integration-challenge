package router

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alcmoraes/go-data-integration-challenge/utils"
	"github.com/stretchr/testify/assert"

	log "github.com/sirupsen/logrus"
)

func TestCompanyApi(t *testing.T) {

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

	t.Run("Create a new company", func(t *testing.T) {

		w := httptest.NewRecorder()

		body := []byte(`{"name": "Facetruck Corp", "zip": "12345", "persist": true}`)

		req, err := http.NewRequest("POST", "/companies/", bytes.NewBuffer(body))
		assert.NoError(t, err)
		R.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Returns a company when it finds one", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", fmt.Sprintf("/companies/?name=%s&zip=%d", "facetruck", 12345), nil)
		assert.NoError(t, err)
		R.ServeHTTP(w, req)
		log.Print(w.Body)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// It looks like goroutines do not get executed when unit testing (?)
	t.Run("Success when uploading CSV", func(t *testing.T) {

		values := map[string]io.Reader{
			"file":    utils.MustOpen("../test/mock/q2_clientData.csv"), // lets assume its this file
			"persist": strings.NewReader("false"),
		}

		w := httptest.NewRecorder()

		h, m, err := utils.Upload("/companies/upload", values)
		h.Header.Set("Content-Type", m.FormDataContentType())

		R.ServeHTTP(w, h)

		log.Print(w.Code)

		assert.NoError(t, err)

	})

}
