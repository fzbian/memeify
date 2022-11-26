package services

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/labstack/gommon/log"
	"meme-generator/entities"
	"meme-generator/enums"
	"meme-generator/interfaces/services"
)

var (
	ABC = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type memeServices struct {
}

func NewMemeServices() services.MemeServices {
	return &memeServices{}
}

func (controller *memeServices) GenerateMeme(nameMeme string, queryParams url.Values) (string, error) {
	configMeme := entities.MemeConfig{}
	if err := bindMemeConfig(&configMeme, nameMeme, queryParams); err != nil {
		return "", err
	}

	im, err := gg.LoadPNG(configMeme.MemePath)
	if err != nil {
		log.Info(err)
		return "", err
	}

	if err = setDraw(im, configMeme); err != nil {
		return "", err
	}

	return configMeme.NameFile, nil
}

func bindMemeConfig(config *entities.MemeConfig, nameMeme string, qValues url.Values) error {
	newMemeConfig, ok := entities.NewMemeConfig[nameMeme]
	if !ok {
		return errors.New("meme name not found")
	}

	config.Name = nameMeme
	config.NameFile = fmt.Sprintf(enums.NameFile, nameMeme)
	config.MemePath = fmt.Sprintf(enums.NamePath, nameMeme)
	config.Color = newMemeConfig.Color
	config.Font = newMemeConfig.Font

	switch {
	case nameMeme == enums.Trump:
		newText := qValues["text"]
		if len(newText) < 1 {
			newText = append(newText, "")
		}
		config.MemeOptions = append(config.MemeOptions, newMemeConfig.MemeOptions[0])
		config.MemeOptions[0].Text = append(config.MemeOptions[0].Text, newText[0])
	case nameMeme == enums.GrimReaperKnockingDoor:
		for i := 0; i < 4; i++ {
			config.MemeOptions = append(config.MemeOptions, newMemeConfig.MemeOptions[i])
			newUrl := qValues["url"+strconv.Itoa(i+1)]
			if len(newUrl) < 1 {
				newUrl = append(newUrl, "")
			}
			config.MemeOptions[i].Url = append(config.MemeOptions[i].Url, newUrl[0])
		}
	}

	return nil
}

func setDraw(im image.Image, config entities.MemeConfig) error {
	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(config.Color)

	if !(config.Name == enums.GrimReaperKnockingDoor) {
		if err := dc.LoadFontFace(config.Font.Path, config.Font.FontSize); err != nil {
			return fmt.Errorf("an error occurred when load font, %v", err)
		}
	}

	switch {
	case config.Name == enums.Trump && len(config.MemeOptions) == enums.TrumpMaxParams:
		dc.DrawImage(im, 0, 0)
		dc.DrawStringWrapped(config.MemeOptions[0].Text[0], config.MemeOptions[0].X, config.MemeOptions[0].Y,
			config.MemeOptions[0].AX, config.MemeOptions[0].AY, config.MemeOptions[0].Width,
			config.MemeOptions[0].LineSpacing, config.MemeOptions[0].Align)
	case config.Name == enums.GrimReaperKnockingDoor && len(config.MemeOptions) == enums.GrimReaperKnockingDoorMaxParams:
		dc.DrawImage(im, 0, 0)
		for _, option := range config.MemeOptions {
			_, img := getImageURL(option.Url[0])
			resizedImg := imaging.Resize(img, option.ResizeWH, option.ResizeWH, imaging.Lanczos)
			dc.DrawImage(resizedImg, int(option.X), int(option.Y))
		}
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

func getImageURL(url string) (string, image.Image) {
	if url == "" {
		url = "https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png"
	}

	response, _ := http.Get(url)

	m, _, err := image.Decode(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	buff := new(bytes.Buffer)
	_ = jpeg.Encode(buff, m, nil)

	nameFile := randomName(15) + ".png"

	//os.WriteFile(nameFile, buff.Bytes(), 0644)

	return nameFile, m
}

func randomName(lenght int) string {
	b := make([]rune, lenght)

	for i := range b {
		b[i] = ABC[rand.Intn(len(ABC))]
	}

	return string(b)
}
