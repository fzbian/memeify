package entities

import (
	"image/color"

	"github.com/fogleman/gg"
)

var (
	NewMemeConfig = map[string]MemeConfig{
		"trump":                     Trump,
		"change_my_mind":            ChangeMyMind,
		"disappointed_black_guy":    DisappointedBlackGuy,
		"distracted_boyfriend":      DistractedBoyfriend,
		"drake":                     Drake,
		"expanding_brain":           ExpandingBrain,
		"jason_momoa_henry_cavil":   JasonMomoaHenryCavil,
		"left_right":                LeftRight,
		"running_away_balloon":      RunningAwayBalloon,
		"spiderman":                 Spiderman,
		"this_is":                   ThisIs,
		"three_headed_dragon":       ThreeHeadedDragon,
		"undertaker":                Undertaker,
		"grim_reaper_knocking_door": GrimReaperKnockingDoor,
		"zoolander":                 Zoolander,
	}

	ChangeMyMind = MemeConfig{
		Font: Font{
			Path:     "fonts/arial.ttf",
			FontSize: 18,
		},
		Color: color.Black,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           200,
					Y:           240,
					AX:          0,
					AY:          0,
					Rotate:      -5,
					Width:       200,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
		},
	}

	DisappointedBlackGuy = MemeConfig{
		Font: Font{
			Path:     "fonts/arial.ttf",
			FontSize: 50,
		},
		Color: color.Black,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           30,
					Y:           30,
					AX:          0,
					AY:          0,
					Width:       680,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           30,
					Y:           400,
					AX:          0,
					AY:          0,
					Width:       680,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
		},
	}

	DistractedBoyfriend = MemeConfig{
		Font: Font{
			Path:     "fonts/impact.ttf",
			FontSize: 50,
		},
		Color: color.Black,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           200,
					Y:           580,
					AX:          0,
					AY:          0,
					Width:       355,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           650,
					Y:           350,
					AX:          0,
					AY:          0,
					Width:       355,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           900,
					Y:           450,
					AX:          0,
					AY:          0,
					Width:       355,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
		},
	}

	Drake = MemeConfig{
		Font: Font{
			Path:     "fonts/arial.ttf",
			FontSize: 50,
		},
		Color: color.Black,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           618,
					Y:           16,
					AX:          0,
					AY:          0,
					Width:       1000,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           618,
					Y:           618,
					AX:          0,
					AY:          0,
					Width:       1000,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
		},
	}

	ExpandingBrain = MemeConfig{
		Font: Font{
			Path:     "fonts/impact.ttf",
			FontSize: 40,
		},
		Color: color.Black,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           10,
					Y:           10,
					AX:          0,
					AY:          0,
					Width:       400,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           10,
					Y:           315,
					AX:          0,
					AY:          0,
					Width:       400,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           10,
					Y:           620,
					AX:          0,
					AY:          0,
					Width:       400,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           10,
					Y:           900,
					AX:          0,
					AY:          0,
					Width:       400,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
		},
	}

	JasonMomoaHenryCavil = MemeConfig{
		Font: Font{
			Path:     "fonts/impact.ttf",
			FontSize: 30,
		},
		Color: color.White,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           44,
					Y:           347,
					AX:          0,
					AY:          0,
					Width:       150,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           220,
					Y:           347,
					AX:          0,
					AY:          0,
					Width:       250,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
		},
	}

	LeftRight = MemeConfig{
		Font: Font{
			Path:     "fonts/impact.ttf",
			FontSize: 24,
		},
		Color: color.White,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           195,
					Y:           100,
					AX:          0,
					AY:          0,
					Width:       130,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           421,
					Y:           100,
					AX:          0,
					AY:          0,
					Width:       130,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
		},
	}

	RunningAwayBalloon = MemeConfig{
		Font: Font{
			Path:     "fonts/impact.ttf",
			FontSize: 30,
		},
		Color: color.Black,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           492,
					Y:           68,
					AX:          0,
					AY:          0,
					Width:       400,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           6,
					Y:           600,
					AX:          0,
					AY:          0,
					Width:       200,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
		},
	}

	Spiderman = MemeConfig{
		Font: Font{
			Path:     "fonts/impact.ttf",
			FontSize: 30,
		},
		Color: color.White,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           158,
					Y:           147,
					AX:          0,
					AY:          0,
					Width:       400,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           550,
					Y:           110,
					AX:          0,
					AY:          0,
					Width:       400,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
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

	Undertaker = MemeConfig{
		Font: Font{
			Path:     "fonts/impact.ttf",
			FontSize: 45,
		},
		Color: color.White,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           220,
					Y:           580,
					AX:          0,
					AY:          0,
					Width:       280,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           536,
					Y:           440,
					AX:          0,
					AY:          0,
					Width:       280,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
		},
	}

	ThreeHeadedDragon = MemeConfig{
		Font: Font{
			Path:     "fonts/impact.ttf",
			FontSize: 45,
		},
		Color: color.White,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           5,
					Y:           333,
					AX:          0,
					AY:          0,
					Width:       170,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           261,
					Y:           291,
					AX:          0,
					AY:          0,
					Width:       280,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           505,
					Y:           340,
					AX:          0,
					AY:          0,
					Width:       170,
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
					X: 35,
					Y: 70,
				},
				Resize: ResizeImage{
					Width:  80,
					Height: 80,
				},
			},
			{
				UrlImg: "",
				DrawImgP: DrawImagePoint{
					X: 225,
					Y: 50,
				},
				Resize: ResizeImage{
					Width:  120,
					Height: 120,
				},
			},
			{
				UrlImg: "",
				DrawImgP: DrawImagePoint{
					X: 520,
					Y: 50,
				},
				Resize: ResizeImage{
					Width:  130,
					Height: 130,
				},
			},
			{
				UrlImg: "",
				DrawImgP: DrawImagePoint{
					X: 400,
					Y: 230,
				},
				Resize: ResizeImage{
					Width:  100,
					Height: 100,
				},
			},
		},
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

	Zoolander = MemeConfig{
		Font: Font{
			Path:     "fonts/arial.ttf",
			FontSize: 50,
		},
		Color: color.White,
		MemeOptions: []MemeOptions{
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           50,
					Y:           530,
					AX:          0,
					AY:          0,
					Width:       500,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
			{
				Text: "",
				DrawStrWrappedP: DrawStringWrappedPoints{
					X:           720,
					Y:           480,
					AX:          0,
					AY:          0,
					Width:       500,
					LineSpacing: 1.5,
					Align:       gg.AlignLeft,
				},
			},
		},
	}
)
