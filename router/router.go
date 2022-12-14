package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, path=${path}, status=${status} latency=${latency_human}\n",
	}))

	basePath := r.server.Group("/api")
	r.server.Static("/", "docs")
	basePath.GET("/health", controllers.HealthCheck)
	r.memeGroup.Resources(basePath)
}
