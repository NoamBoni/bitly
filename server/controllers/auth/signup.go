package auth

import (
	"github.com/NoamBoni/bitly/server/helpers"
	"github.com/NoamBoni/bitly/server/models"
	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		helpers.SendError(ctx, err, 400)
		return
	}
	if err := user.Insert(); err != nil {
		helpers.SendError(ctx, err, 400)
		return
	}
	if token, err := GenerateJWT(user.Id); err != nil {
		helpers.SendError(ctx, err, 400)
		return
	} else {
		SignCookie(ctx,"token", token)
		ctx.JSON(200, gin.H{"data": user})
	}
}
