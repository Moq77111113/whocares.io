package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/moq77111113/whocares/config"
	"github.com/moq77111113/whocares/pkg/context"
)

// Stores the configuration in the request context, making it accessible to ui.
func Config(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set(context.ConfigKey, cfg)
			return next(ctx)
		}
	}
}
