package providers

import (
	"go.uber.org/dig"
	"meme-generator/controllers"
	"meme-generator/router"
	"meme-generator/router/groups"
	"meme-generator/server"
	"meme-generator/services"
	"meme-generator/utils"
)

var (
	Container *dig.Container
)

func BuildContainer() *dig.Container {
	Container = dig.New()
    
	_ = Container.Provide(server.NewServer)
	_ = Container.Provide(router.NewRouter)
	_ = Container.Provide(groups.NewMemeGroup)
	_ = Container.Provide(controllers.NewMemeController)
	_ = Container.Provide(services.NewMemeServices)
	_ = Container.Provide(utils.NewUtils)
    
	return Container
}
