package services

import (
	"bytes"
	"github.com/fogleman/gg"
	"image/jpeg"
	"log"
	"meme-generator/entities"
	"meme-generator/interfaces/services"
	"os"
)

type memeServices struct {
}

func NewMemeServices() services.MemeServices {
	return &memeServices{}
}

func (controller *memeServices) GenerateMeme(config entities.MemeConfig) string {
	im, err := gg.LoadPNG(config.MemePath)
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(config.Color)
	dc.LoadFontFace(config.Font.Path, config.Font.FontSize)
	dc.DrawImage(im, 0, 0)
	options := config.MemeOptions[0]
	dc.DrawStringWrapped(options.Text, options.X, options.Y, options.AX, options.AY, options.Width, options.LineSpacing, options.Align)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	nameFile := config.NameFile
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}
