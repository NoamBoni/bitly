package urls

import (
	"github.com/NoamBoni/bitly/server/helpers"
	"github.com/NoamBoni/bitly/server/models"
	"github.com/gin-gonic/gin"
)

func UpdateURL(ctx *gin.Context) {
	var reqUrl models.Url
	if err := ctx.ShouldBindJSON(&reqUrl); err != nil {
		helpers.SendError(ctx, err, 400)
		return
	}
	if err := reqUrl.Update(); err != nil {
		helpers.SendError(ctx, err, 400)
		return
	}
	ctx.JSON(201, gin.H{"data": reqUrl})
}
