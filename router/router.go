package router

import (
	"meme-generator/controllers"
	"meme-generator/interfaces/router/groups"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		Format: "_________\nMethod: ${method}\nPath: ${path}\nStatus code: ${status}\nLatency: ${laency_human}\nRemote IP: ${remote_ip}\nHost: ${host}\nProtocol: ${protocol}\nReferer: ${referer}\n",
	}))

	basePath := r.server.Group("/api")
	r.server.Static("/", "docs")
	basePath.GET("/health", controllers.HealthCheck)
	r.memeGroup.Resources(basePath)
}
