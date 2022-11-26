package groups

import (
	"github.com/labstack/echo/v4"
	"meme-generator/interfaces/controllers"
	"meme-generator/interfaces/router/groups"
)

type MemesGroup struct {
	memeController controllers.MemeController
}

func NewMemeGroup(memeController controllers.MemeController) groups.MemeGroup {
	return &MemesGroup{memeController}
}

func (routes MemesGroup) Resources(group *echo.Group) {
	groupPath := group.Group("")
	groupPath.GET("/:nameMeme", routes.memeController.GenerateMeme)
}
