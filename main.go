package main

import (
	"fmt"
	"net/http"
	"os"

	"meme-generator/functions/images"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	basePath := e.Group("/api")

	// Custom Meme
	basePath.GET("/custom/", customMeme)

	// Template Memes
	basePath.GET("/trump/", trumpMeme)                                  // Trump
	basePath.GET("/change_my_mind/", changeMyMind)                      // Change My Mind
	basePath.GET("/disappointed_black_guy/", disappointedBlackGuy)      // Disappointed Black Guy
	basePath.GET("/distracted_boyfriend/", distractedBoyfriend)         // Disappointed Black Guy
	basePath.GET("/drake/", drake)                                      // Drake
	basePath.GET("/expanding_brain/", expandingBrain)                   // Expanding Brain
	basePath.GET("/jason_momoa_henry_cavil/", jasonMomoaHenryCavil)     // Jason Momoa Henry Cavil
	basePath.GET("/left_right/", leftRight)                             // Left Right
	basePath.GET("/running_away_balloon/", runningAwayBalloon)          // Running Away Balloon
	basePath.GET("/spiderman/", spiderman)                              // Spiderman
	basePath.GET("/three_headed_dragon/", threeHeadedDragon)            // Three Headed Dragon
	basePath.GET("/undertaker/", undertaker)                            // Undertaker
	basePath.GET("/this_is/", thisIs)                                   // This Is
	basePath.GET("/grim_reaper_knocking_door/", grimReaperKnockingDoor) // Grim Reaper Knocking Door

	basePath.GET("/", index)
	e.Logger.Fatal(e.Start(":8080"))
}

func customMeme(c echo.Context) error {
	fmt.Println("CustomMeme Getting...")
	textTop := c.QueryParam("top")
	textBottom := c.QueryParam("bottom")
	url := c.QueryParam("url")

	newMeme := images.CustomMeme(textTop, textBottom, url)

	defer os.Remove(newMeme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(newMeme, newMeme)
}

func trumpMeme(c echo.Context) error {
	text := c.QueryParam("text")
	trump := images.TrumpMeme(text)
	defer os.Remove(trump)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(trump, trump)
}

func changeMyMind(c echo.Context) error {
	text := c.QueryParam("text")
	meme := images.ChangeMyMind(text)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func disappointedBlackGuy(c echo.Context) error {
	text1 := c.QueryParam("text1")
	text2 := c.QueryParam("text2")
	meme := images.DisappointedBlackGuy(text1, text2)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func distractedBoyfriend(c echo.Context) error {
	text1 := c.QueryParam("text1")
	text2 := c.QueryParam("text2")
	text3 := c.QueryParam("text3")
	meme := images.DistractedBoyfriend(text1, text2, text3)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func drake(c echo.Context) error {
	text1 := c.QueryParam("text1")
	text2 := c.QueryParam("text2")
	meme := images.Drake(text1, text2)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func expandingBrain(c echo.Context) error {
	text1 := c.QueryParam("text1")
	text2 := c.QueryParam("text2")
	text3 := c.QueryParam("text3")
	text4 := c.QueryParam("text4")
	meme := images.ExpandingBrain(text1, text2, text3, text4)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func jasonMomoaHenryCavil(c echo.Context) error {
	text1 := c.QueryParam("text1")
	text2 := c.QueryParam("text2")
	meme := images.JasonMomoaHenryCavil(text1, text2)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func leftRight(c echo.Context) error {
	text1 := c.QueryParam("text1")
	text2 := c.QueryParam("text2")
	meme := images.LeftRight(text1, text2)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func runningAwayBalloon(c echo.Context) error {
	text1 := c.QueryParam("text1")
	text2 := c.QueryParam("text2")
	text3 := c.QueryParam("text3")
	meme := images.RunningAwayBalloon(text1, text2, text3)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func spiderman(c echo.Context) error {
	text1 := c.QueryParam("text1")
	text2 := c.QueryParam("text2")
	meme := images.Spiderman(text1, text2)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func threeHeadedDragon(c echo.Context) error {
	text1 := c.QueryParam("text1")
	text2 := c.QueryParam("text2")
	text3 := c.QueryParam("text3")
	meme := images.ThreeHeadedDragon(text1, text2, text3)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func undertaker(c echo.Context) error {
	text1 := c.QueryParam("text1")
	text2 := c.QueryParam("text2")
	meme := images.Undertaker(text1, text2)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func thisIs(c echo.Context) error {
	url := c.QueryParam("url")
	meme := images.ThisIs(url)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func grimReaperKnockingDoor(c echo.Context) error {
	url1 := c.QueryParam("imageURL1")
	url2 := c.QueryParam("imageURL2")
	url3 := c.QueryParam("imageURL3")
	url4 := c.QueryParam("imageURL4")
	meme := images.GrimReaperKnockingDoor(url1, url2, url3, url4)
	defer os.Remove(meme)
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline(meme, meme)
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Ok!")
}
