package main

import (
	"github.com/alcmoraes/go-data-integration-challenge/config"
	"github.com/alcmoraes/go-data-integration-challenge/router"
)

func main() {
	gin := router.Load()
	gin.Run(":" + config.Config.Port)
}
