package handlers

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/moq77111113/whocares/internal/services"
	"github.com/moq77111113/whocares/internal/ui"
	"github.com/moq77111113/whocares/internal/ui/models"
	"github.com/moq77111113/whocares/internal/ui/pages"
	"github.com/moq77111113/whocares/pkg/og"

	"github.com/labstack/echo/v4"
)

type Pages struct {
	counter  *services.Counter
	messages *services.Messages
	og       *og.Generator
}

func init() {
	Register(new(Pages))
}

func (h *Pages) Init(c *services.Container) error {
	h.counter = c.Counter
	h.messages = c.Message
	h.og = c.OG
	return nil
}

func (h *Pages) Routes(g *echo.Group) {
	g.GET("/", h.home).Name = "home"
	g.GET("/counter", h.count).Name = "counter"
}

func (h *Pages) home(c echo.Context) error {
	data, err := h.buildPageData(c, true)
	if err != nil {
		return err
	}

	r := ui.NewRequest(c)

	var buf strings.Builder
	_ = pages.Home(r, data).Render(&buf)
	return c.HTML(http.StatusOK, buf.String())
}

func (h *Pages) count(c echo.Context) error {
	data, err := h.buildPageData(c, false)
	if err != nil {
		return err
	}

	var buf strings.Builder
	_ = pages.CounterContent(data.Count, data.Message, data.Subtext).Render(&buf)
	return c.HTML(http.StatusOK, buf.String())
}

// buildPageData constructs all the data needed for page rendering
func (h *Pages) buildPageData(c echo.Context, includeOG bool) (*models.Counter, error) {
	data := &models.Counter{
		Count:  h.counter.GetCount(),
		Target: c.QueryParam("target"),
	}

	messageSet, err := h.loadMessages()
	if err != nil {
		return nil, err
	}

	data.Message = h.selectRandomMessage(messageSet.Primary)
	data.Subtext = h.selectRandomMessage(messageSet.Secondary)

	if includeOG {
		ogImageURL, err := h.generateOGImage(data)
		if err != nil {
			return nil, err
		}
		data.OgImageURL = ogImageURL
	}

	return data, nil
}

// loadMessages loads the default message variant
func (h *Pages) loadMessages() (*services.MessageSet, error) {
	return h.messages.LoadVariant(services.VariantDefault)
}

// selectRandomMessage picks a random message from the provided slice
func (h *Pages) selectRandomMessage(messages []string) string {
	if len(messages) == 0 {
		return ""
	}
	return messages[rand.Intn(len(messages))]
}

// generateOGImage creates the Open Graph image for the given page data
func (h *Pages) generateOGImage(data *models.Counter) (string, error) {
	path, err := h.og.Generate(data, og.ThemeBrutalist)
	if err != nil {
		return "", err
	}
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	return path, nil
}
