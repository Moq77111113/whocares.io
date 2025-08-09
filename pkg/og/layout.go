package og

import (
	"fmt"
	"strings"

	"github.com/fogleman/gg"
)

const (
	CounterYPosition     = 0.25
	MessageYPosition     = 0.65
	MaxCounterWidth      = 1
	LineHeight           = 60.0
	BrandBottomMargin    = 20.0
	BrandRightMargin     = 30.0
	LineHeightMultiplier = 0.8
)

type DrawConfig struct {
	Width     int
	Height    int
	Fonts     FontSet
	Theme     Theme
	BrandText string
}

func DrawBackground(dc *gg.Context, config DrawConfig) {
	dc.SetColor(config.Theme.Background)
	dc.DrawRectangle(0, 0, float64(config.Width), float64(config.Height))
	dc.Fill()
}

func DrawCounter(dc *gg.Context, count string, config DrawConfig) {
	centerY := float64(config.Height) * CounterYPosition

	dc.SetColor(config.Theme.PrimaryText)

	if config.Fonts.Huge != nil {
		dc.SetFontFace(config.Fonts.Huge)
	}

	textWidth, textHeight := dc.MeasureString(count)

	maxWidth := float64(config.Width) * MaxCounterWidth

	fmt.Printf("Counter text width: %f, max width: %f\n", textWidth, maxWidth)
	if textWidth > maxWidth {
		dc.SetColor(config.Theme.PrimaryText)
		if config.Fonts.Large != nil {
			dc.SetFontFace(config.Fonts.Large)
		}
		textWidth, textHeight = dc.MeasureString(count)
	}

	x := (float64(config.Width) - textWidth) / 2
	y := centerY + textHeight/2

	dc.DrawString(count, x, y)
}

func DrawMessage(dc *gg.Context, message string, config DrawConfig) {
	centerY := float64(config.Height) * MessageYPosition

	dc.SetColor(config.Theme.PrimaryText)
	if config.Fonts.Large != nil {
		dc.SetFontFace(config.Fonts.Large)
	}

	wrappedText := WrapText(message, MaxTextWidth)
	lines := strings.Split(wrappedText, "\n")

	totalHeight := float64(len(lines)) * LineHeight
	startY := centerY - totalHeight/2

	for i, line := range lines {
		textWidth, _ := dc.MeasureString(line)
		x := (float64(config.Width) - textWidth) / 2
		y := startY + float64(i)*LineHeight + LineHeight*LineHeightMultiplier

		dc.DrawString(line, x, y)
	}
}

func DrawBrand(dc *gg.Context, config DrawConfig) {
	if config.BrandText == "" {
		return
	}

	dc.SetColor(config.Theme.AccentText)
	if config.Fonts.Small != nil {
		dc.SetFontFace(config.Fonts.Small)
	}

	textWidth, _ := dc.MeasureString(config.BrandText)
	x := float64(config.Width) - textWidth - BrandRightMargin
	y := float64(config.Height) - BrandBottomMargin

	dc.DrawString(config.BrandText, x, y)
}
