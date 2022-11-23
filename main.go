package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"meme-generator/functions/files"
	"meme-generator/functions/images"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	basePath := e.Group("/api")

	// Old Memes
	basePath.GET("/old/:meme/:top/:bottom", OldMeme)

	// New Template Memes
	basePath.GET("/trump/:text", trumpMeme) // Trump

	// Meme by URL and Params
	basePath.GET("/byurl/:text/:x/:y/:url", imageByURL)

	basePath.GET("/", index)
	e.Logger.Fatal(e.Start(":8080"))
}

func OldMeme(c echo.Context) error {
	fmt.Println("localMeme Getting...")
	textTop := c.Param("top")
	textBottom := c.Param("bottom")
	nameMeme := c.Param("meme")

	newMeme := images.CreateOldMeme(nameMeme, textTop, textBottom)

	defer os.Remove(newMeme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(newMeme, newMeme)
}

func imageByURL(c echo.Context) error {
	fmt.Println("imageByURL Getting...")
	urlImg := c.Param("url")
	text := c.Param("text")
	x, _ := strconv.ParseFloat(c.Param("x"), 64)
	y, _ := strconv.ParseFloat(c.Param("y"), 64)

	newMeme := images.CreateMemeByURL(urlImg, text, x, y)
	defer os.Remove(newMeme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(newMeme, newMeme)
}

func trumpMeme(c echo.Context) error {
	text := c.Param("text")
	trump := images.TrumpMeme(text)
	defer os.Remove(trump)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(trump, trump)
}

func index(c echo.Context) error {
	fmt.Println("index Getting...")
	return c.JSON(http.StatusOK, files.GlobalPaths)
}
