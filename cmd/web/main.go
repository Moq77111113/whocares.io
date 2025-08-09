package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/moq77111113/whocares/internal/handlers"
	"github.com/moq77111113/whocares/internal/services"
)

func main() {

	c := services.NewContainer()
	defer func() {
		if err := c.Close(); err != nil {
			log.Fatalf("Error closing container: %v", err)
			os.Exit(1)
		}
	}()

	if err := handlers.BuildRoutes(c); err != nil {
		log.Fatalf("Error building routes: %v", err)
		os.Exit(1)
	}

	go func() {
		srv := http.Server{
			Addr:         fmt.Sprintf("%s:%d", c.Config.Server.Host, c.Config.Server.Port),
			Handler:      c.Web,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		}

		if err := c.Web.StartServer(&srv); errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("shutting down the server: %v", err)
			os.Exit(1)
		}
		fmt.Println("Server started on", srv.Addr)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
}
