package controllers

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"meme-generator/entities"
	"meme-generator/interfaces/controllers"
	"meme-generator/interfaces/services"
)

type memeController struct {
	services services.MemeServices
}

func NewMemeController(services services.MemeServices) controllers.MemeController {
	return &memeController{services}
}

func (controller *memeController) GenerateMeme(ctx echo.Context) error {
	filter := entities.FilterMeme{}
	if err := ctx.Bind(&filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	generatedMeme, err := controller.services.GenerateMeme(filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	defer func() {
		_ = os.Remove(generatedMeme)
	}()

	ctx.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return ctx.Inline(generatedMeme, generatedMeme)
}
