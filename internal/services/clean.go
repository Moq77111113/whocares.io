package services

import (
	"os"
	"path/filepath"
	"time"

	"github.com/moq77111113/whocares/config"
)

const (
	ogSubPath = "og"
)

type Cleaner struct {
	config *config.Config
}

func NewCleaner(cfg *config.Config) *Cleaner {
	return &Cleaner{
		config: cfg,
	}
}

func (s *Cleaner) CleanOldImages() error {
	ogDir := filepath.Join(s.config.Static.PublicDir, ogSubPath)

	return filepath.Walk(ogDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && time.Since(info.ModTime()) > time.Duration(s.config.App.CacheDuration) {
			return os.Remove(path)
		}

		return nil
	})
}
