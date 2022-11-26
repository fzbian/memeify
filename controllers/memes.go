package controllers

import (
	"github.com/labstack/echo/v4"
	"meme-generator/entities"
	"meme-generator/interfaces/controllers"
	"meme-generator/interfaces/services"
	"net/http"
)

type memeController struct {
	services services.MemeServices
}

func NewMemeController(services services.MemeServices) controllers.MemeController {
	return &memeController{services}
}

func (controller *memeController) GenerateMeme(ctx echo.Context) error {
	nameMeme := ctx.Param("nameMeme")
	text := ctx.QueryParam("text")
	memeConfig, ok := entities.NewMemeConfig[nameMeme]
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "meme name not found")
	}
	memeConfig.MemeOptions[0].Text = append(memeConfig.MemeOptions[0].Text, text)
	trump := controller.services.GenerateMeme(nameMeme, memeConfig)

	//defer func() {
	//	os.Remove(trump)
	//}()

	ctx.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return ctx.Inline(trump, trump)
}
