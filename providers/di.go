package providers

import (
	"go.uber.org/dig"
	"meme-generator/router"
	"meme-generator/server"
)

var (
	Container *dig.Container
)

func BuildContainer() *dig.Container {
	Container = dig.New()
	_ = Container.Provide(server.NewServer)
	_ = Container.Provide(router.NewRouter)
	return Container
}
