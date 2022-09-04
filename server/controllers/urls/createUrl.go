package urls

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/NoamBoni/bitly/server/helpers"
	"github.com/NoamBoni/bitly/server/models"
	"github.com/gin-gonic/gin"
)

const length = 8

func CreateURL(ctx *gin.Context) {
	userId, exists := ctx.Get("user-id")
	id := userId.(int)
	if !exists {
		helpers.SendError(ctx, errors.New("something's wrong, try again later"), 500)
		return
	}
	var newUrl models.Url
	if err := ctx.ShouldBindJSON(&newUrl); err != nil {
		helpers.SendError(ctx, err, 400)
		return
	}
	newUrl.User_id = id
	newUrl.Modified_url = createRandomString()
	if err := newUrl.Insert(); err != nil {
		helpers.SendError(ctx, err, 500)
		return
	}
	if urls, err := models.GetUrlsByUserId(id); err != nil {
		helpers.SendError(ctx, err, 500)
		return
	} else {
		ctx.JSON(201, gin.H{"data": urls})
	}
}

func createRandomString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
