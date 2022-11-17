package main

import (
	"net/http"
	"os"

	"meme-generator/functions/files"
	"meme-generator/functions/images"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	basePath := e.Group("/api")
	basePath.GET("/:meme/:top/:bottom", localMeme)
	basePath.GET("/", index)
	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, files.GlobalPaths)
}

func localMeme(c echo.Context) error {
	textTop := c.Param("top")
	textBottom := c.Param("bottom")
	nameMeme := c.Param("meme")

	newMeme := images.CreateMeme(nameMeme, textTop, textBottom)

	defer os.Remove(newMeme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(newMeme, newMeme)
}
