package routes

import (
	"github.com/NoamBoni/bitly/server/controllers"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine{
	router := gin.Default()
	router.Use(gin.CustomRecovery(controllers.RecoveryMiddleware))
	addAuthRoutes(router)
	router.Use(controllers.AuthMiddleware)
	addURLRoutes(router)
	
	return router
}