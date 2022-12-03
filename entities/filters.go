package entities

import (
	"context"

	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)

var (
	Validate = validator.New()
	Conform  = modifiers.New()
)

type FilterMeme struct {
	Name  string `param:"nameMeme"`
	Text  string `query:"text" mod:"ucase"`
	Text1 string `query:"text1" mod:"ucase"`
	Text2 string `query:"text2" mod:"ucase"`
	Text3 string `query:"text3" mod:"ucase"`
	Text4 string `query:"text4" mod:"ucase"`
	Url   string `query:"url" mod:"default=https://raw.githubusercontent.com/fzbian/meme-generator/main/memes/Insert_image_here.png" validate:"url"`
	Url1  string `query:"url1" mod:"default=https://raw.githubusercontent.com/fzbian/meme-generator/main/memes/Insert_image_here.png" validate:"url"`
	Url2  string `query:"url2" mod:"default=https://raw.githubusercontent.com/fzbian/meme-generator/main/memes/Insert_image_here.png" validate:"url"`
	Url3  string `query:"url3" mod:"default=https://raw.githubusercontent.com/fzbian/meme-generator/main/memes/Insert_image_here.png" validate:"url"`
	Url4  string `query:"url4" mod:"default=https://raw.githubusercontent.com/fzbian/meme-generator/main/memes/Insert_image_here.png" validate:"url"`
}

func (f *FilterMeme) Validate() error {
	if err := Conform.Struct(context.Background(), f); err != nil {
		return err
	}
	return Validate.Struct(f)
}
