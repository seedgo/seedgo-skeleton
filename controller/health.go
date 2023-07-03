package controller

import (
	"github.com/gin-gonic/gin"
	"seedgo-skeleton/response"
)

func Health(ctx *gin.Context) {
	response.Success(ctx, "ok")
}
