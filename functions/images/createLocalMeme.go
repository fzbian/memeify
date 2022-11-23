package images

import (
	"bytes"
	"fmt"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"meme-generator/functions/files"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

func CreateOldMeme(imagePath string, textTop string, textBottom string) string {
	im := files.GetFile(imagePath)
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
	nameFile := fmt.Sprintf("%s.jpg", imagePath)
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func TrumpMeme(text string) string {
	im, err := gg.LoadPNG("memes/trump.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.LoadFontFace("fonts/arial.ttf", 50)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(text, 750, 580, 0, 0, 430, 1.5, gg.AlignLeft)
	dc.Clip()
	dc.SavePNG("trump.png")

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "trump.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}
