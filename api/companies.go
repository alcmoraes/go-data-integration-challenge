package api

import (
	"net/http"

	"github.com/alcmoraes/go-data-integration-challenge/database"
	"github.com/alcmoraes/go-data-integration-challenge/importer"
	"github.com/alcmoraes/go-data-integration-challenge/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

func GetCompany(c *gin.Context) {

	var json types.Company

	if err := c.ShouldBindQuery(&json); err == nil {
		session, collection, err := database.GetCollection("companies")
		defer session.Close()

		if err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
			return
		}

		company := types.Company{}
		err = collection.Find(bson.M{"name": bson.M{"$regex": "^" + json.Name + "\\s?.*", "$options": "i"}, "zip": json.Zip}).One(&company)

		if err != nil {
			log.Error(err)
			if err.Error() == "not found" {
				c.JSON(http.StatusNotFound, gin.H{"status": "ERROR", "message": err.Error()})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, company)
		return
	} else {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}

}

func ImportCompany(c *gin.Context) {

	var json types.Company

	if err := c.ShouldBindWith(&json, binding.JSON); err == nil {

		if json.Website == "" && !json.Persist {
			c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": "Website shouldn't be null"})
			return
		}

		if err := database.AddCompanyIntoDatabase(json, json.Persist); err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Done!"})
		}

	} else {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
	}

}

func UploadCompany(c *gin.Context) {
	if file, handler, err := c.Request.FormFile("file"); err == nil {
		defer file.Close()

		var body types.CompanyUploadQuery

		err := c.ShouldBindWith(&body, binding.FormMultipart)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
			return
		}

		doneImporting := make(chan bool, 1)

		go importer.Worker(file, handler, body.Persist, doneImporting)

		c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Done!"})

	} else {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
	}
}
