package og

import (
	"path/filepath"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

const (
	HugeFontSize   = 150.0
	LargeFontSize  = 48.0
	MediumFontSize = 36.0
	SmallFontSize  = 24.0

	MinCounterFontSize = 120.0
	FontSizeStep       = 20.0
)

type FontSet struct {
	Huge   font.Face
	Large  font.Face
	Medium font.Face
	Small  font.Face
}

type FontConfig struct {
	FontsDir  string
	Bold      string
	ExtraBold string
	Regular   string
}

func NewFontConfig(fontsDir string) FontConfig {
	// TODO: Remove this once we have a proper config system
	return FontConfig{
		FontsDir:  fontsDir,
		Bold:      "JetBrainsMono-Bold.ttf",
		ExtraBold: "JetBrainsMono-ExtraBold.ttf",
		Regular:   "JetBrainsMono-Regular.ttf",
	}
}

func (fc FontConfig) BoldPath() string {
	return filepath.Join(fc.FontsDir, fc.Bold)
}

func (fc FontConfig) ExtraBoldPath() string {
	return filepath.Join(fc.FontsDir, fc.ExtraBold)
}

func (fc FontConfig) RegularPath() string {
	return filepath.Join(fc.FontsDir, fc.Regular)
}

func loadFont(path string, size float64, fallbackPath string) font.Face {

	if face, err := gg.LoadFontFace(path, size); err == nil {
		return face
	}
	if face, err := gg.LoadFontFace(fallbackPath, size); err == nil {
		return face
	}

	return nil
}

func LoadFonts(config FontConfig) FontSet {
	fonts := FontSet{
		Huge:   loadFont(config.ExtraBoldPath(), HugeFontSize, config.BoldPath()),
		Large:  loadFont(config.RegularPath(), LargeFontSize, config.BoldPath()),
		Medium: loadFont(config.BoldPath(), MediumFontSize, config.RegularPath()),
		Small:  loadFont(config.RegularPath(), SmallFontSize, config.BoldPath()),
	}

	if !areFontsLoaded(fonts) {
		fallback := basicfont.Face7x13
		return FontSet{
			Huge:   fallback,
			Large:  fallback,
			Medium: fallback,
			Small:  fallback,
		}
	}

	return fonts
}

func areFontsLoaded(fonts FontSet) bool {
	return fonts.Huge != nil || fonts.Large != nil || fonts.Medium != nil || fonts.Small != nil
}
