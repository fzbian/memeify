package main

import (
	"meme-generator/providers"
	"meme-generator/router"

	"github.com/labstack/echo/v4"
)

func main() {
	container := providers.BuildContainer()

	err := container.Invoke(func(server *echo.Echo, router *router.Router) {
		router.Init()
		server.Logger.Fatal(server.Start(":8080"))
	})

	if err != nil {
		panic(err)
	}
}
