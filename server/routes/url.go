package routes

import (
	"github.com/NoamBoni/bitly/server/controllers/urls"
	"github.com/gin-gonic/gin"
)

func addURLRoutes(router *gin.Engine){
	urlRoutes := router.Group("/url")
	{
		urlRoutes.POST("/",urls.CreateURL)
		urlRoutes.PUT("/",urls.UpdateURL)
	}
}