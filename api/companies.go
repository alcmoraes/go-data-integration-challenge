package api

import (
	"net/http"

	"github.com/alcmoraes/go-data-integration-challenge/database"
	"github.com/alcmoraes/go-data-integration-challenge/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/globalsign/mgo/bson"
)

func GetCompany(c *gin.Context) {

	var json types.Company

	if err := c.ShouldBindQuery(&json); err == nil {
		session, collection, err := database.GetCollection("companies")
		defer session.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		company := types.Company{}
		err = collection.Find(bson.M{"name": bson.M{"$regex": "^" + json.Name + "\\s?.*", "$options": "i"}, "zip": json.Zip}).One(&company)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, company)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func MergeCompany(c *gin.Context) {

	var json types.Company

	if err := c.ShouldBindWith(&json, binding.JSON); err == nil {

		if json.Website == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Website should not be null"})
			return
		}

		session, collection, err := database.GetCollection("companies")
		defer session.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		companies := types.Companies{}
		err = collection.Find(bson.M{"name": bson.M{"$regex": "^" + json.Name + "\\s?.*", "$options": "i"}, "zip": json.Zip}).All(&companies.Companies)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		if len(companies.Companies) > 1 {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Multiple companies match the given data"})
			return
		}

		if collection.UpdateId(companies.Companies[0].ID, json); err == nil {
			c.JSON(http.StatusOK, bson.M{"status": "OK"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
