package router

import (
	"github.com/alcmoraes/go-data-integration-challenge/api"
	"github.com/gin-gonic/gin"
)

// Load the gin HTTP
func Load() *gin.Engine {
	ginRouter := gin.Default()

	companies := ginRouter.Group("/companies")
	{
		companies.GET("/", api.GetCompany)
		companies.POST("/", api.ImportCompany)
		companies.POST("/upload", api.UploadCompany)
	}

	return ginRouter
}
