package entities

import (
	"github.com/fogleman/gg"
	"image/color"
)

type MemeConfig struct {
	Name        string
	MemePath    string
	NameFile    string
	Font        Font
	Color       color.Color
	MemeOptions []MemeOptions
}

type Font struct {
	Path     string
	FontSize float64
}

type MemeOptions struct {
	Text        string
	X           float64
	Y           float64
	AX          float64
	AY          float64
	Width       float64
	LineSpacing float64
	Align       gg.Align
}

var (
	Trump = MemeConfig{
		Name:     "trump",
		MemePath: "memes/trump.png",
		NameFile: "trump.png",
		Font: Font{
			Path:     "fonts/arial.ttf",
			FontSize: 50,
		},
		Color: color.Black,
		MemeOptions: []MemeOptions{
			{
				X:           750,
				Y:           580,
				AX:          0,
				AY:          0,
				Width:       430,
				LineSpacing: 1.5,
				Align:       gg.AlignLeft,
			},
		},
	}
)
