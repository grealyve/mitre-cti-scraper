package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grealyve/mitre-cti-scraper/controller"
)

func AptFeedRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/apt_feed", controller.GetAptFeed)
}
