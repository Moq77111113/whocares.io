package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/moq77111113/whocares/internal/services"
)

var handlers []Handler

type Handler interface {
	Routes(g *echo.Group)

	Init(*services.Container) error
}

func Register(h Handler) {
	handlers = append(handlers, h)
}

func Get() []Handler {
	return handlers
}
