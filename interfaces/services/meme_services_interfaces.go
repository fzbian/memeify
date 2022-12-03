package services

import (
	"meme-generator/entities"
)

type MemeServices interface {
	GenerateMeme(filter entities.FilterMeme) (string, error)
}
