package auth

import (
	"os"
	"time"

	"github.com/NoamBoni/bitly/server/helpers"
	"github.com/gin-gonic/gin"
)

func SignCookie(ctx *gin.Context, key, value string) {
	helpers.LoadEnv()
	secure := os.Getenv("SERVER_STATE") != "development"
	ctx.SetCookie(key, value, int(time.Hour)*3, "/", "localhost", secure, secure)
}
