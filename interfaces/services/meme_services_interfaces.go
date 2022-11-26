package services

import "meme-generator/entities"

type MemeServices interface {
	GenerateMeme(config entities.MemeConfig) string
}
