package services

import (
	"net/url"

	"meme-generator/entities"
	"meme-generator/interfaces/services"
	"meme-generator/interfaces/utils"

	"github.com/labstack/gommon/log"
)

type memeServices struct {
	utils utils.Utils
}

func NewMemeServices(utils utils.Utils) services.MemeServices {
	return &memeServices{utils}
}

func (services *memeServices) GenerateMeme(nameMeme string, queryParams url.Values) (string, error) {
	configMeme := entities.MemeConfig{}
	if err := services.utils.BindMemeConfig(&configMeme, nameMeme, queryParams); err != nil {
		return "", err
	}

	img, err := services.utils.LoadPNG(configMeme.MemePath)
	if err != nil {
		log.Info(err)
		return "", err
	}

	if err = services.utils.DrawMeme(img, configMeme); err != nil {
		return "", err
	}

	return configMeme.NameFile, nil
}
