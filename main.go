package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"math"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"github.com/labstack/echo/v4"
)

func getFile(nameFile string) image.Image {
	path := map[string]string{
		"bad-luck-brian":      "memes/bad-luck-brian.jpg",
		"batman-slaps-robin":  "memes/batman-slaps-robin.jpg",
		"crying-peter-parker": "memes/crying-peter-parker.jpg",
		"filosoraptor":        "memes/filosoraptor.jpg",
		"gopher":              "memes/gopher.jpg",
		"kermit-the-frog":     "memes/kermit-the-frog.jpg",
		"mr-bean":             "memes/mr-bean.jpg",
		"trump":               "memes/trump.jpg",
		"willy-wonka":         "memes/willy-wonka.jpg",
	}
	f, err := os.Open(path[nameFile])
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	im, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return im
}

func createMeme(imagePath string, textTop string, textBottom string) {
	im := getFile(imagePath)
	bounds := im.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	max := math.Round(float64(width) / 10)

	dc := gg.NewContextForImage(im)

	if err := dc.LoadFontFace("fonts/impact.ttf", max); err != nil {
		panic(err)
	}

	positionX := float64(width / 2)
	positionTopY := float64(height / 6)
	positionBottomY := float64(5 * height / 6)

	dc.SetRGB(0, 0, 0)
	n := 6
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				continue
			}
			x := positionX + float64(dx)
			ytop := positionTopY + float64(dy)
			ybottom := positionBottomY + float64(dy)
			dc.DrawStringAnchored(strings.ToUpper(textTop), x, ytop, 0.5, 0)
			dc.DrawStringAnchored(strings.ToUpper(textBottom), x, ybottom, 0.5, 1)
		}
	}

	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(strings.ToUpper(textTop), positionX, positionTopY, 0.5, 0)
	dc.DrawStringAnchored(strings.ToUpper(textBottom), positionX, positionBottomY, 0.5, 1)

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	os.WriteFile("out.jpg", buff.Bytes(), 0644)
}

func main() {
	e := echo.New()
	e.GET("/:meme/:top/:bottom", handler)
	e.Logger.Fatal(e.Start(":1323"))
}

func handler(c echo.Context) error {
	textTop := c.Param("top")
	textBottom := c.Param("bottom")
	nameMeme := c.Param("meme")

	createMeme(nameMeme, textTop, textBottom)

	defer os.Remove("out.jpg")
	c.Response().Header().Set(echo.HeaderContentType, "image/jpeg")
	return c.Inline("out.jpg", "out.jpg")
}
