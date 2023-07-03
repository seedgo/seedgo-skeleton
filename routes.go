package main

import (
	"github.com/gin-gonic/gin"
	"seedgo-skeleton/controller"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", controller.Health)
}
