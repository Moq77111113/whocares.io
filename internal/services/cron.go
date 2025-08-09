package services

import (
	"log"
	"time"
)

type CronService struct {
	imageService *Cleaner
	ticker       *time.Ticker
	done         chan bool
}

func NewCronService(imageService *Cleaner) *CronService {
	return &CronService{
		imageService: imageService,
		done:         make(chan bool),
	}
}

func (c *CronService) Start() {
	c.ticker = time.NewTicker(10 * time.Minute)

	go func() {
		for {
			select {
			case <-c.ticker.C:
				if err := c.imageService.CleanOldImages(); err != nil {
					log.Printf("Error cleaning old images: %v", err)
				} else {
					log.Println("Successfully cleaned old OG images")
				}
			case <-c.done:
				return
			}
		}
	}()

	log.Println("🧹 Cron service started - cleaning OG images every 10 minutes")
}

func (c *CronService) Stop() {
	if c.ticker != nil {
		c.ticker.Stop()
	}
	c.done <- true
	log.Println("🛑 Cron service stopped")
}
