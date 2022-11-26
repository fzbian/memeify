package services

import (
	"bytes"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/labstack/gommon/log"
	"image"
	"image/jpeg"
	"meme-generator/entities"
	"meme-generator/enums"
	"meme-generator/interfaces/services"
	"os"
)

const (
	TrumpMaxParams                  = 1
	GrimReaperKnockingDoorMaxParams = 4
)

type memeServices struct {
}

func NewMemeServices() services.MemeServices {
	return &memeServices{}
}

func (controller *memeServices) GenerateMeme(nameMeme string, config entities.MemeConfig) string {
	name, _ := entities.NameMemes[nameMeme]
	setNameMemeConfig(name, &config)
	im, err := gg.LoadPNG(config.MemePath)
	if err != nil {
		log.Info(err)
		return ""
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(config.Color)
	dc.LoadFontFace(config.Font.Path, config.Font.FontSize)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringWrapped(config.MemeOptions[0].Text[0], config.MemeOptions[0].X, config.MemeOptions[0].Y,
		config.MemeOptions[0].AX, config.MemeOptions[0].AY, config.MemeOptions[0].Width,
		config.MemeOptions[0].LineSpacing, config.MemeOptions[0].Align)
	dc.Clip()

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	os.WriteFile(config.NameFile, buff.Bytes(), 0644)
	return config.NameFile
}

func setNameMemeConfig(name string, config *entities.MemeConfig) {
	config.Name = name
	config.NameFile = fmt.Sprintf(enums.NameFile, name)
	config.MemePath = fmt.Sprintf(enums.NamePath, name)
}

func setDraw(name string, im image.Image, config entities.MemeConfig) {
	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(config.Color)
	dc.LoadFontFace(config.Font.Path, config.Font.FontSize)
	dc.DrawImage(im, 0, 0)
	switch {
	case name == "trump" && len(config.MemeOptions) == TrumpMaxParams:
		dc.DrawStringWrapped(config.MemeOptions[0].Text[0], config.MemeOptions[0].X, config.MemeOptions[0].Y,
			config.MemeOptions[0].AX, config.MemeOptions[0].AY, config.MemeOptions[0].Width,
			config.MemeOptions[0].LineSpacing, config.MemeOptions[0].Align)
		dc.Clip()
	case name == "grim_reaper_knocking_door" && len(config.MemeOptions) == GrimReaperKnockingDoorMaxParams:
		for _, option := range config.MemeOptions {
			dc.DrawStringWrapped(option.Text[0], option.X, option.Y, option.AX, option.AY, option.Width, option.LineSpacing, option.Align)
		}
		dc.Clip()
	}
	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	os.WriteFile(config.NameFile, buff.Bytes(), 0644)
}
