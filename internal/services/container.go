package services

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/moq77111113/whocares/config"
	"github.com/moq77111113/whocares/pkg/og"
)

// Container contains all the services used in the application
type Container struct {
	Web *echo.Echo

	Config *config.Config

	Cleanup *Cleaner

	Message *Messages

	Counter *Counter

	OG *og.Generator
}

func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	c.initWeb()
	c.initCleanup()
	c.initMessage()
	c.initCounter()
	c.initGenerator()
	return c
}

func (c *Container) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := c.Web.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
func (c *Container) initConfig() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	c.Config = cfg
}

func (c *Container) initWeb() {
	c.Web = echo.New()
	c.Web.HideBanner = true

	c.Web.Static("/", c.Config.Static.PublicDir)
}

func (c *Container) initCleanup() {
	c.Cleanup = NewCleaner(c.Config)
}

func (c *Container) initMessage() {
	c.Message = NewMessageService(c.Config)
}

func (c *Container) initCounter() {
	c.Counter = NewCounter(c.Config)
}

func (c *Container) initGenerator() {
	c.OG = og.NewGenerator(c.Config, fmt.Sprintf("%s/og", c.Config.Static.PublicDir))
}
