package assets

import "embed"

//go:embed messages/*.yaml
var Messages embed.FS
