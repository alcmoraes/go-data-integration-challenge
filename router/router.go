package router

import (
	"github.com/alcmoraes/go-data-integration-challenge/api"
	"github.com/gin-gonic/gin"
)

// Load the gin HTTP
func Load() *gin.Engine {
	router := gin.Default()

	companies := router.Group("/companies")
	{
		companies.GET("/", api.GetCompany)
		companies.POST("/", api.MergeCompany)
	}

	return router
}
