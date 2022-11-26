package groups

import (
	"github.com/labstack/echo/v4"
	"meme-generator/functions/images"
	"meme-generator/interfaces/controllers"
	"meme-generator/interfaces/router/groups"
	"os"
)

type MemesGroup struct {
	memeController controllers.MemeController
}

func NewMemeGroup(memeController controllers.MemeController) groups.MemeGroup {
	return &MemesGroup{memeController}
}

func (routes MemesGroup) Resources(group *echo.Group) {
	groupPath := group.Group("")
	groupPath.GET("/:nameMeme/", routes.memeController.GenerateMeme)
}

func TrumpMeme(c echo.Context) error {
	text := c.QueryParam("text")
	trump := images.TrumpMeme(text)
	defer os.Remove(trump)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(trump, trump)
}
