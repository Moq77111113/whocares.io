package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/moq77111113/whocares/config"
	"github.com/moq77111113/whocares/pkg/context"
	g "maragu.dev/gomponents"
)

type (
	// Request encapsulates information needed by UI components and layouts.
	Request struct {
		// Title is the page title.
		Title string

		// Context stores the Echo context.
		Context echo.Context

		// CurrentPath is the path of the current request.
		CurrentPath string

		// Metatags stores standard meta tag values.
		Metatags struct {
			Description string
			Keywords    []string
		}

		// OpenGraph stores Open Graph/Twitter card values.
		OpenGraph struct {
			Title       string
			Description string
			ImageURL    string
		}

		// Config holds the application configuration.
		Config *config.Config
	}

	// LayoutFunc renders content wrapped in a layout for the given request.
	LayoutFunc func(*Request, g.Node) g.Node
)

// NewRequest generates a new Request using the Echo context of a given HTTP request.
func NewRequest(ctx echo.Context) *Request {
	r := &Request{
		Context:     ctx,
		CurrentPath: ctx.Request().URL.Path,
	}

	if cfg := ctx.Get(context.ConfigKey); cfg != nil {
		r.Config = cfg.(*config.Config)
	}

	return r
}
