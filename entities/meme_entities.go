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
	Rotate      float64
}

type ResizeImage struct {
	Width  int
	Height int
}
