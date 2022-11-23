package images

import (
	"bytes"
	"image/color"
	"image/jpeg"
	"meme-generator/functions/files"
	"os"

	"github.com/fogleman/gg"
)

func CreateMemeByURL(url string, text string, x, y float64) string {
	nameFile, im := files.GetImageURL(url)

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.LoadFontFace("fonts/arial.ttf", 50)
	dc.DrawImage(im, 0, 0)
	dc.DrawString(text, x, y)
	dc.Clip()
	dc.SavePNG(nameFile)

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}
