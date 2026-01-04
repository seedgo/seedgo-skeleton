package main

import (
	"seedgo-skeleton/controller"

	"github.com/seedgo/seedgo"
)

func main() {
	server := seedgo.NewServer()

	controller.RegisterRoutes(server.GetEngine())

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
