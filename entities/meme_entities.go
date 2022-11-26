package entities

import (
	"image/color"

	"github.com/fogleman/gg"
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
	ResizeWH    int
}

var (
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
				X:        32,
				Y:        68,
				ResizeWH: 50,
			},
			{
				X:        153,
				Y:        57,
				ResizeWH: 80,
			},
			{
				X:        355,
				Y:        20,
				ResizeWH: 90,
			},
			{
				X:        280,
				Y:        150,
				ResizeWH: 50,
			},
		},
	}
)
