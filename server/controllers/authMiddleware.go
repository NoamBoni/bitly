package controllers

import (
	"errors"
	"os"
	"time"

	"github.com/NoamBoni/bitly/server/controllers/auth"
	"github.com/NoamBoni/bitly/server/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(ctx *gin.Context) {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		helpers.SendError(ctx,err,403)
		return
	}
	token, err := jwt.ParseWithClaims(
		cookie,
		&auth.Claims{},
		getSecret,
	)
	if err != nil {
		helpers.SendError(ctx, errors.New("please login to continue"), 403)
		return
	}
	claims, ok := token.Claims.(*auth.Claims)
	if !ok {
		helpers.SendError(ctx, errors.New("please login to continue"), 403)
		return
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		helpers.SendError(ctx, errors.New("please login again"), 403)
		return
	}
	ctx.Set("user-id", claims.Id)
	ctx.Next()
}

func getSecret(t *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}
