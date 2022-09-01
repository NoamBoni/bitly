package helpers

import (

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, err error, code int){
	ctx.JSON(code, gin.H{"error": err.Error()})
}