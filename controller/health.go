package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/seedgo/seedgo"
)

func Health(ctx *gin.Context) {
	seedgo.Success(ctx, "ok")
}
