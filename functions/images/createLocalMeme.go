package images

import (
	"bytes"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"meme-generator/functions/files"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

func CustomMeme(top string, bottom string, url string) string {
	nameFile, im := files.GetImageURL(url)
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
			dc.DrawStringAnchored(strings.ToUpper(top), x, ytop, 0.5, 0)
			dc.DrawStringAnchored(strings.ToUpper(bottom), x, ybottom, 0.5, 1)
		}
	}

	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(strings.ToUpper(top), positionX, positionTopY, 0.5, 0)
	dc.DrawStringAnchored(strings.ToUpper(bottom), positionX, positionBottomY, 0.5, 1)

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
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

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "trump.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func ChangeMyMind(text string) string {
	im, err := gg.LoadPNG("memes/change_my_mind.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.LoadFontFace("fonts/impact.ttf", 18)
	dc.DrawImage(im, 0, 0)
	dc.Rotate(gg.Radians(-5))
	dc.DrawStringWrapped(text, 200, 240, 0, 0, 200, 1.5, gg.AlignLeft)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "change_my_mind.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func DisappointedBlackGuy(text1 string, text2 string) string {
	im, err := gg.LoadPNG("memes/disappointed_black_guy.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.LoadFontFace("fonts/impact.ttf", 50)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(text1, 30, 30, 0, 0, 680, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text2, 30, 400, 0, 0, 680, 1.5, gg.AlignLeft)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "disappointed_black_guy.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func DistractedBoyfriend(text1 string, text2 string, text3 string) string {
	im, err := gg.LoadPNG("memes/distracted_boyfriend.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.LoadFontFace("fonts/impact.ttf", 50)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(text1, 200, 580, 0, 0, 355, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text2, 650, 350, 0, 0, 355, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text3, 900, 450, 0, 0, 355, 1.5, gg.AlignLeft)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "distracted_boyfriend.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func Drake(text1 string, text2 string) string {
	im, err := gg.LoadPNG("memes/drake.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.LoadFontFace("fonts/impact.ttf", 80)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(text1, 618, 16, 0, 0, 1000, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text2, 618, 618, 0, 0, 1000, 1.5, gg.AlignLeft)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "drake.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func ExpandingBrain(text1 string, text2 string, text3 string, text4 string) string {
	im, err := gg.LoadPNG("memes/expanding_brain.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.LoadFontFace("fonts/impact.ttf", 40)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(text1, 10, 10, 0, 0, 400, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text2, 10, 315, 0, 0, 400, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text3, 10, 620, 0, 0, 400, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text4, 10, 900, 0, 0, 400, 1.5, gg.AlignLeft)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "drake.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func JasonMomoaHenryCavil(text1 string, text2 string) string {
	im, err := gg.LoadPNG("memes/jason_momoa_henry_cavil.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.White)
	dc.LoadFontFace("fonts/impact.ttf", 30)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(text1, 44, 347, 0, 0, 150, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text2, 220, 347, 0, 0, 150, 1.5, gg.AlignLeft)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "drake.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func LeftRight(text1 string, text2 string) string {
	im, err := gg.LoadPNG("memes/left_right.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.White)
	dc.LoadFontFace("fonts/impact.ttf", 24)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(text1, 195, 100, 0, 0, 130, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text2, 421, 100, 0, 0, 130, 1.5, gg.AlignLeft)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "drake.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func RunningAwayBalloon(text1 string, text2 string, text3 string) string {
	im, err := gg.LoadPNG("memes/running_away_balloon.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.LoadFontFace("fonts/impact.ttf", 30)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(text1, 492, 68, 0, 0, 400, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text2, 600, 600, 0, 0, 400, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text3, 6, 750, 0, 0, 200, 1.5, gg.AlignLeft)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "drake.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func Spiderman(text1 string, text2 string) string {
	im, err := gg.LoadPNG("memes/spiderman.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.LoadFontFace("fonts/impact.ttf", 30)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(text1, 158, 147, 0, 0, 400, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text2, 550, 110, 0, 0, 400, 1.5, gg.AlignLeft)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "drake.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func ThreeHeadedDragon(text1 string, text2 string, text3 string) string {
	im, err := gg.LoadPNG("memes/three_headed_dragon.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.White)
	dc.LoadFontFace("fonts/impact.ttf", 30)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(text1, 5, 333, 0, 0, 170, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text2, 261, 291, 0, 0, 170, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text3, 505, 340, 0, 0, 170, 1.5, gg.AlignLeft)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "drake.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func Undertaker(text1 string, text2 string) string {
	im, err := gg.LoadPNG("memes/undertaker.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.White)
	dc.LoadFontFace("fonts/impact.ttf", 45)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(text1, 220, 580, 0, 0, 280, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(text2, 536, 440, 0, 0, 280, 1.5, gg.AlignLeft)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "drake.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}

func ThisIs(imageURL string) string {
	im, err := gg.LoadPNG("memes/this_is.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.White)
	dc.LoadFontFace("fonts/impact.ttf", 150)
	_, newIm := files.GetImageURL(imageURL)
	dc.DrawImage(im, 0, 0)
	dc.DrawImage(newIm, 1100, 150)
	dc.DrawString("Esto es basura", 300, 1260)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := "drake.jpg"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}
