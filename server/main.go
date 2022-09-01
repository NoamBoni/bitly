package main

import (
	"github.com/NoamBoni/bitly/server/helpers"
	"github.com/NoamBoni/bitly/server/models"
	"github.com/NoamBoni/bitly/server/routes"
)

func main() {
	helpers.LoadEnv()
	defer models.Db.Close()
	router := routes.CreateRouter()
	router.Run(":8000")
}
