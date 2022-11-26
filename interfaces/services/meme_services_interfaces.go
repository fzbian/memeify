package services

import "meme-generator/entities"

type MemeServices interface {
	GenerateMeme(nameMeme string, config entities.MemeConfig) string
}
