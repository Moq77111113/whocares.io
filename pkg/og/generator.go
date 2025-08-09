package og

import (
	"fmt"
	"path/filepath"

	"github.com/fogleman/gg"
	"github.com/moq77111113/whocares/config"
	"github.com/moq77111113/whocares/internal/ui/models"
)

const (
	ImageWidth  = 1200
	ImageHeight = 630
)

type Generator struct {
	cfg       *config.Config
	ogDir     string
	fonts     FontSet
	fontCfg   FontConfig
	brandText string
}

func NewGenerator(cfg *config.Config, ogDir string) *Generator {
	fontCfg := NewFontConfig(cfg.Static.FontsDir)

	return &Generator{
		cfg:       cfg,
		fonts:     LoadFonts(fontCfg),
		fontCfg:   fontCfg,
		ogDir:     ogDir,
		brandText: cfg.Base.Title,
	}
}

func (g *Generator) Generate(c *models.Counter, themeName ThemeName) (string, error) {
	theme := GetTheme(&themeName)
	cacheKey := GenerateCacheKey(c.Count, c.Message, c.Target)

	if cachedPath, exists := CheckCache(cacheKey, g.ogDir); exists {
		return cachedPath, nil
	}

	dc := gg.NewContext(ImageWidth, ImageHeight)

	config := DrawConfig{
		Width:     ImageWidth,
		Height:    ImageHeight,
		Fonts:     g.fonts,
		Theme:     theme,
		BrandText: g.brandText,
	}

	DrawBackground(dc, config)
	DrawCounter(dc, c.Count, config)
	DrawMessage(dc, c.Message, config)
	DrawBrand(dc, config)

	filename := fmt.Sprintf("%s%s", cacheKey, FileExtension)
	outputPath := filepath.Join(g.ogDir, filename)

	if err := dc.SavePNG(outputPath); err != nil {
		return "", err
	}

	return outputPath, nil
}
