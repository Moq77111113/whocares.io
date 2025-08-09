package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/moq77111113/whocares/config"
)

type Counter struct {
	cfg       *config.Config
	baseCount int
}

func NewCounter(cfg *config.Config) *Counter {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return &Counter{
		cfg:       cfg,
		baseCount: cfg.App.Seed,
	}
}

func (c *Counter) GetCount() string {

	variation := rand.Intn(500000)
	count := c.baseCount + variation

	return formatNumber(count)
}

func formatNumber(n int) string {
	str := fmt.Sprintf("%d", n)
	result := ""

	for i, char := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += ","
		}
		result += string(char)
	}

	return result
}
