package utils

import (
	"image"
	"net/url"

	"meme-generator/entities"
)

type Utils interface {
	BindMemeConfig(config *entities.MemeConfig, nameMeme string, qValues url.Values) error
	LoadPNG(path string) (image.Image, error)
	DrawMeme(imageMeme image.Image, config entities.MemeConfig) error
}
