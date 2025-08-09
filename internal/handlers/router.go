package handlers

import (
	"net/http"
	"time"

	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/moq77111113/whocares/internal/services"
	"github.com/moq77111113/whocares/pkg/middleware"
)

func BuildRoutes(c *services.Container) error {

	c.Web.Group("").Static("/public", c.Config.Static.PublicDir)

	g := c.Web.Group("")

	g.Use(echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}),
		echomw.Recover(),
		echomw.Logger(),
		echomw.RequestID(),
		echomw.Gzip(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: 30 * time.Second,
		}),
		middleware.Config(c.Config),
	)

	for _, h := range Get() {
		if err := h.Init(c); err != nil {
			return err
		}
		h.Routes(g)
	}

	return nil
}
