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
	Url         []string
	Text        []string
	X           float64
	Y           float64
	AX          float64
	AY          float64
	Width       float64
	LineSpacing float64
	Align       gg.Align
}

var (
	NameMemes = map[string]string{
		"trump":                     "trump",
		"grim_reaper_knocking_door": "grim_reaper_knocking_door",
	}

	NewMemeConfig = map[string]MemeConfig{
		"trump":                     Trump,
		"grim_reaper_knocking_door": GrimReaperKnockingDoor,
	}

	Trump = MemeConfig{
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
	GrimReaperKnockingDoor = MemeConfig{
		Name:     "grim_reaper_knocking_door",
		MemePath: "memes/grim_reaper_knocking_door.png",
		NameFile: "grim_reaper_knocking_door.png",
		Color:    color.White,
		MemeOptions: []MemeOptions{
			{
				X: 32,
				Y: 68,
			},
			{
				X: 153,
				Y: 57,
			},
			{
				X: 355,
				Y: 20,
			},
			{
				X: 280,
				Y: 150,
			},
		},
	}
)

func (mc MemeConfig) Validate() error {
	return nil
}
