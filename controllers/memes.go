package controllers

import (
	"github.com/labstack/echo/v4"
	"meme-generator/entities"
	"meme-generator/interfaces/controllers"
	"meme-generator/interfaces/services"
	"os"
)

type memeController struct {
	services services.MemeServices
}

func NewMemeController(services services.MemeServices) controllers.MemeController {
	return &memeController{services}
}

func (controller *memeController) GenerateMeme(ctx echo.Context) error {
	text := ctx.QueryParam("text")
	trumpConfig := entities.Trump
	trumpConfig.MemeOptions[0].Text = text
	trump := controller.services.GenerateMeme(trumpConfig)
	defer os.Remove(trump)
	ctx.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return ctx.Inline(trump, trump)
}
