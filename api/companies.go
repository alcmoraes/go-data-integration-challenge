package api

import (
	"net/http"

	"github.com/alcmoraes/go-data-integration-challenge/database"
	"github.com/alcmoraes/go-data-integration-challenge/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/mgo.v2/bson"
)

func GetCompany(c *gin.Context) {

	var json types.Company

	if err := c.ShouldBindWith(&json, binding.JSON); err == nil {
		session, collection, err := database.GetCollection("companies")
		defer session.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		company := types.Company{}
		err = collection.Find(bson.M{"name": json.Name, "zip": json.Zip}).One(&company)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, company)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
