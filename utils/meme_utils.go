package utils

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/goombaio/namegenerator"
	"github.com/labstack/gommon/log"
	"meme-generator/entities"
	"meme-generator/enums"
	interfaces "meme-generator/interfaces/utils"
)

type utils struct {
	randomName namegenerator.Generator
}

func NewUtils(randomName namegenerator.Generator) interfaces.Utils {
	return &utils{randomName}
}

func (u utils) BindMemeConfig(config *entities.MemeConfig, nameMeme string, qValues url.Values) error {
	newMemeConfig, ok := entities.NewMemeConfig[nameMeme]
	if !ok {
		return errors.New("meme name not found")
	}

	*config = newMemeConfig
	config.Name = nameMeme
	config.NameFile = fmt.Sprintf(enums.NameFile, nameMeme)
	config.MemePath = fmt.Sprintf(enums.NamePath, nameMeme)

	switch {
	case nameMeme == enums.Trump:
		newText, exist := qValues["text"]
		u.assignTextOrUrl(&config.MemeOptions[0].Text, exist, newText)
		break
	case nameMeme == enums.GrimReaperKnockingDoor:
		for i := 0; i < 4; i++ {
			newUrl, exist := qValues["url"+strconv.Itoa(i+1)]
			u.assignTextOrUrl(&config.MemeOptions[i].UrlImg, exist, newUrl)
		}
		break
	case nameMeme == enums.ThisIs:
		newUrl, exist := qValues["url"]
		u.assignTextOrUrl(&config.MemeOptions[0].UrlImg, exist, newUrl)
		break
	}

	return nil
}

func (u utils) assignTextOrUrl(text *string, exist bool, textOrUrlToAssign []string) {
	if !exist {
		*text = enums.EmptyText
		return
	}

	*text = textOrUrlToAssign[0]
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
	dc := gg.NewContextForImage(img)
	dc.Clear()
	dc.SetColor(config.Color)
	dc.DrawImage(img, 0, 0)

	if !(config.Name == enums.GrimReaperKnockingDoor) {
		if err := dc.LoadFontFace(config.Font.Path, config.Font.FontSize); err != nil {
			return fmt.Errorf("an error occurred when load font, %v", err)
		}
	}

	switch {
	case config.Name == enums.Trump:
		option := config.MemeOptions[0]
		dc.DrawStringWrapped(option.Text, option.DrawStrWrappedP.X, option.DrawStrWrappedP.Y,
			option.DrawStrWrappedP.AX, option.DrawStrWrappedP.AY, option.DrawStrWrappedP.Width,
			option.DrawStrWrappedP.LineSpacing, option.DrawStrWrappedP.Align)
	case config.Name == enums.GrimReaperKnockingDoor:
		for _, option := range config.MemeOptions {
			_, img, err := u.getImageURL(option.UrlImg)
			if err != nil {
				return err
			}
			resizedImg := imaging.Resize(img, option.Resize.Width, option.Resize.Width, imaging.Lanczos)
			dc.DrawImage(resizedImg, option.DrawImgP.X, option.DrawImgP.Y)
		}
	case config.Name == enums.ThisIs:
		option := config.MemeOptions[0]
		dc.DrawString(option.Text, option.DrawStrP.X, option.DrawStrP.Y)
		_, img, err := u.getImageURL(config.MemeOptions[0].UrlImg)
		if err != nil {
			return err
		}
		resizedImg := imaging.Resize(img, option.Resize.Width, option.Resize.Height, imaging.Lanczos)
		dc.DrawImage(resizedImg, option.DrawImgP.X, option.DrawImgP.Y)
	}

	dc.Clip()
	buff := new(bytes.Buffer)

	if err := jpeg.Encode(buff, dc.Image(), nil); err != nil {
		return fmt.Errorf("an error occurred when encode image, %v", err)
	}

	if err := os.WriteFile(config.NameFile, buff.Bytes(), 0644); err != nil {
		return fmt.Errorf("an error ocurred when write file, %v", err)
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
