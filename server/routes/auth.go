package routes

import (
	"github.com/NoamBoni/bitly/server/controllers/auth"
	"github.com/gin-gonic/gin"
)

func addAuthRoutes(router *gin.Engine){
	authentication := router.Group("/auth")
	{
		authentication.POST("/signup",auth.Signup)
		authentication.POST("/login",auth.Login)
	}
}