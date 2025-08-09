package og

import "image/color"

const (
	DefaultTheme ThemeName = "brutalist"
)

type Theme struct {
	Background    color.Color
	PrimaryText   color.Color
	SecondaryText color.Color
	AccentText    color.Color
}

type ThemeName string

const (
	ThemeBrutalist ThemeName = "brutalist"
)

var themes = map[ThemeName]Theme{
	ThemeBrutalist: {
		Background:    color.RGBA{12, 10, 18, 255},
		PrimaryText:   color.RGBA{243, 248, 240, 255},
		SecondaryText: color.RGBA{164, 183, 160, 255},
		AccentText:    color.RGBA{255, 112, 166, 230},
	},
}

func GetTheme(name *ThemeName) Theme {
	if name == nil {
		return themes[DefaultTheme]
	}

	if theme, exists := themes[*name]; exists {
		return theme
	}

	return themes[DefaultTheme]
}
