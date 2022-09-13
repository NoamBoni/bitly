package auth

import (
	"net/http"

	"github.com/NoamBoni/bitly/server/helpers"
	"github.com/NoamBoni/bitly/server/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginbody struct {
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func Login(ctx *gin.Context) {
	var body loginbody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		helpers.SendError(ctx, err, 400)
		return
	}
	user := models.User{
		Email: body.Email,
	}
	if err := user.SelectByEmail(); err != nil {
		helpers.SendError(ctx, err, 403)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(body.Password)); err != nil {
		helpers.SendError(ctx, err, 403)
		return
	}
	if token, err := GenerateJWT(user.Id); err != nil {
		helpers.SendError(ctx, err, 400)
		return
	} else {
		SignCookie(ctx, "token", token)
		user.Password = nil
		ctx.JSON(http.StatusOK, gin.H{"data": user})
	}
}
