package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware(c *gin.Context, recovered interface{}){
if err, ok := recovered.(string); ok {
      c.String(500, fmt.Sprintf("error: %s", err))
    }
    c.AbortWithStatus(500)
}