package main

import (
	"net/http"

	"github.com/NoamBoni/bitly/server/helpers"
	"github.com/NoamBoni/bitly/server/models"
	"github.com/gin-gonic/gin"
)

func main() {
	helpers.LoadEnv()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		user := models.User{
			Firstname: "noam",
			Lastname:  "boni",
			Email:     "e@io.io",
			Password:  "123456",
		}
		user.Insert()
		c.JSON(http.StatusOK, gin.H{
			"message": "noam",
			"user":    user,
		})
	})
	r.Run()
}
