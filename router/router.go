package router

import (
	"github.com/labstack/echo/v4"
	"meme-generator/controllers"
	"meme-generator/interfaces/router/groups"
)

type Router struct {
	server    *echo.Echo
	memeGroup groups.MemeGroup
}

func NewRouter(server *echo.Echo, group groups.MemeGroup) *Router {
	return &Router{server, group}
}

func (r *Router) Init() {
	basePath := r.server.Group("/api")

	basePath.GET("/health", controllers.HealthCheck)
	r.memeGroup.Resources(basePath)
}
