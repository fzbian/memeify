package utils

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"

	"meme-generator/entities"
	"meme-generator/enums"
	interfaces "meme-generator/interfaces/utils"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/goombaio/namegenerator"
	"github.com/labstack/gommon/log"
)

type utils struct {
	randomName namegenerator.Generator
}

func NewUtils(randomName namegenerator.Generator) interfaces.Utils {
	return &utils{randomName}
}
func (u utils) BindMemeConfig(config *entities.MemeConfig, filter entities.FilterMeme) error {
	newMemeConfig, exist := entities.NewMemeConfig[filter.Name]
	if !exist {
		return fmt.Errorf("meme name '%s' not found", filter.Name)
	}
	*config = newMemeConfig
	config.SetMissingValues(filter)
	return nil
}

func (u utils) LoadPNG(path string) (image.Image, error) {
	im, err := gg.LoadPNG(path)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	return im, nil
}

func (u utils) DrawMeme(img image.Image, config entities.MemeConfig) error {
	newMeme, err := u.drawer(img, config.Name, config)
	if err != nil {
		return err
	}

	buff := new(bytes.Buffer)

	if err := jpeg.Encode(buff, newMeme, nil); err != nil {
		return fmt.Errorf("an error occurred when encode image, %v", err)
	}

	if err := os.WriteFile(config.NameFile, buff.Bytes(), 0644); err != nil {
		return fmt.Errorf("an error ocurred when write file, %v", err)
	}

	return nil
}

func (u utils) drawer(img image.Image, memeName string, config entities.MemeConfig) (image.Image, error) {
	dc := gg.NewContextForImage(img)
	dc.Clear()
	dc.SetColor(config.Color)
	dc.DrawImage(img, 0, 0)

	if !(config.Name == enums.GrimReaperKnockingDoor) {
		if err := dc.LoadFontFace(config.Font.Path, config.Font.FontSize); err != nil {
			return nil, fmt.Errorf("an error occurred when load font, %v", err)
		}
	}

	options := config.MemeOptions

	switch memeName {
	case enums.ChangeMyMind:
		dc.Rotate(gg.Radians(options[0].DrawStrWrappedP.Rotate))
		u.drawStringWrappeds(dc, options)
	case enums.DisappointedBlackGuy:
		u.drawStringWrappeds(dc, options)
	case enums.DistractedBoyfriend:
		u.drawStringWrappeds(dc, options)
	case enums.Drake:
		u.drawStringWrappeds(dc, options)
	case enums.ExpandingBrain:
		u.drawStringWrappeds(dc, options)
	case enums.JasonMomoaHenryCavil:
		u.drawStringWrappeds(dc, options)
	case enums.LeftRight:
		u.drawStringWrappeds(dc, options)
	case enums.RunningAwayBalloon:
		u.drawStringWrappeds(dc, options)
	case enums.Spiderman:
		u.drawStringWrappeds(dc, options)
	case enums.ThisIs:
		u.drawStrings(dc, options)
		err := u.drawImages(dc, options)
		if err != nil {
			return nil, err
		}
	case enums.ThreeHeadedDragon:
		u.drawStringWrappeds(dc, options)
	case enums.Undertaker:
		u.drawStringWrappeds(dc, options)
	case enums.GrimReaperKnockingDoor:
		err := u.drawImages(dc, options)
		if err != nil {
			return nil, err
		}
	case enums.Trump:
		u.drawStringWrappeds(dc, options)
	}

	dc.Clip()

	return dc.Image(), nil
}

func (u utils) drawStringWrappeds(dc *gg.Context, options []entities.MemeOptions) {
	for _, option := range options {
		dc.DrawStringWrapped(option.Text, option.DrawStrWrappedP.X, option.DrawStrWrappedP.Y,
			option.DrawStrWrappedP.AX, option.DrawStrWrappedP.AY, option.DrawStrWrappedP.Width,
			option.DrawStrWrappedP.LineSpacing, option.DrawStrWrappedP.Align)
	}
}

func (u utils) drawStrings(dc *gg.Context, options []entities.MemeOptions) {
	for _, option := range options {
		dc.DrawString(option.Text, option.DrawStrP.X, option.DrawStrP.Y)
	}
}

func (u utils) drawImages(dc *gg.Context, options []entities.MemeOptions) error {
	for _, option := range options {
		_, img, err := u.getImageURL(option.UrlImg)
		if err != nil {
			return err
		}
		resizedImg := imaging.Resize(img, option.Resize.Width, option.Resize.Width, imaging.Lanczos)
		dc.DrawImage(resizedImg, option.DrawImgP.X, option.DrawImgP.Y)
	}
	return nil
}

func (u utils) getImageURL(url string) (string, image.Image, error) {
	if url == "" {
		im, err := gg.LoadPNG("memes/Insert_image_here.png")
		if err != nil {
			log.Info(err)
		}
		return "", im, nil
	}

	response, err := http.Get(url)
	if err != nil {
		return "", nil, errors.New("url invalid")
	}

	m, _, err := image.Decode(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	buff := new(bytes.Buffer)
	_ = jpeg.Encode(buff, m, nil)

	nameFile := u.randomName.Generate() + ".png"

	//os.WriteFile(nameFile, buff.Bytes(), 0644)

	return nameFile, m, nil
}
