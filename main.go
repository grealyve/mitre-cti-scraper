package main

import (
	"github.com/gin-gonic/gin"
	"github.com/grealyve/mitre-cti-scraper/charts"
	"github.com/grealyve/mitre-cti-scraper/middlewares"
	"github.com/grealyve/mitre-cti-scraper/routes"
)

func main() {
	//config.Connect() -> Connect to db, implement when it'll be necessary
	router := gin.Default()
	
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())

	charts.ScetchDiagrams()

	routes.AptFeedRoute(router)

	router.Run("localhost:7777")
}
