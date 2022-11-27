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
	UrlImg          string
	Text            string
	DrawStrP        DrawStringPoints
	DrawImgP        DrawImagePoint
	DrawStrWrappedP DrawStringWrappedPoints
	Resize          ResizeImage
}

type DrawStringPoints struct {
	X float64
	Y float64
}

type DrawImagePoint struct {
	X int
	Y int
}

type DrawStringWrappedPoints struct {
	X           float64
	Y           float64
	AX          float64
	AY          float64
	Width       float64
	LineSpacing float64
	Align       gg.Align
	ResizeWH    int
}

type ResizeImage struct {
	Width  int
	Height int
}

var (
	NewMemeConfig = map[string]MemeConfig{
		"trump":                     Trump,
		"grim_reaper_knocking_door": GrimReaperKnockingDoor,
		"this_is":                   ThisIs,

	}

	Trump = MemeConfig{
		Font: Font{
			Path:     "fonts/arial.ttf",
			FontSize: 50,
		},
		Color: color.Black,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           750,
					Y:           580,
					AX:          0,
					AY:          0,
					Width:       430,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
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
				UrlImg: "",
				DrawImgP: DrawImagePoint{
					X: 32,
					Y: 68,
				},
				Resize: ResizeImage{
					Width:  50,
					Height: 50,
				},
			},
			{
				UrlImg: "",
				DrawImgP: DrawImagePoint{
					X: 153,
					Y: 57,
				},
				Resize: ResizeImage{
					Width:  80,
					Height: 80,
				},
			},
			{
				UrlImg: "",
				DrawImgP: DrawImagePoint{
					X: 355,
					Y: 20,
				},
				Resize: ResizeImage{
					Width:  90,
					Height: 90,
				},
			},
			{
				UrlImg: "",
				DrawImgP: DrawImagePoint{
					X: 280,
					Y: 150,
				},
				Resize: ResizeImage{
					Width:  50,
					Height: 50,
				},
			},
		},
	}

	ThisIs = MemeConfig{
		Font: Font{
			Path:     "fonts/impact.ttf",
			FontSize: 150,
		},
		Color: color.White,
		MemeOptions: []MemeOptions{
			{
				Text:   "Esto es Basura",
				UrlImg: "",
				DrawImgP: DrawImagePoint{
					X: 1050,
					Y: 140,
				},
				DrawStrP: DrawStringPoints{
					X: 300,
					Y: 1260,
				},
				Resize: ResizeImage{
					Width:  500,
					Height: 500,
				},
			}},
	}
)
