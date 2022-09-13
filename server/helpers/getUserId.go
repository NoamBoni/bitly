package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserId(ctx *gin.Context) (int, error){
	userId, exists := ctx.Get("user-id")
	id := userId.(int)
	if !exists {
		return -1, errors.New("something's wrong, try again later")
	}
	return id, nil
}