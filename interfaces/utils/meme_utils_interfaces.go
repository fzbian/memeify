package utils

import (
	"image"

	"meme-generator/entities"
)

type Utils interface {
	LoadPNG(path string) (image.Image, error)
	DrawMeme(imageMeme image.Image, config entities.MemeConfig) error
	BindMemeConfig(config *entities.MemeConfig, filter entities.FilterMeme) error
}
