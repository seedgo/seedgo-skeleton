package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"seedgo-skeleton/common"
	"seedgo-skeleton/config"
	"strconv"
)

func main() {
	config.InitCmd()
	config.Init()

	if !common.ServerConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	RegisterRoutes(r)

	port := strconv.Itoa(common.ServerConfig.Port)
	fmt.Printf("start server at port: %s\n", port)
	panic(r.Run(":" + port))
}
