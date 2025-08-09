package services

import (
	"fmt"
	"html"
	"os"
	"path/filepath"
	"strings"

	"github.com/moq77111113/whocares/assets"
	"github.com/moq77111113/whocares/config"
	"gopkg.in/yaml.v3"
)

type MessageSet struct {
	Primary   []string `yaml:"primary"`
	Secondary []string `yaml:"secondary"`
	Footnotes []string `yaml:"footnote"`
}

type Variant string

const (
	VariantDefault   Variant = "default"
	VariantCorpo     Variant = "corpo"
	VariantSarcastic Variant = "sarcastic"
	VariantWholesome Variant = "wholesome"
)

type Messages struct {
	dir string
}

func NewMessageService(cfg *config.Config) *Messages {
	return &Messages{dir: cfg.Static.MessagesDir}
}

func (c *Messages) LoadVariant(variant Variant) (*MessageSet, error) {
	names := []string{fmt.Sprintf("%s.yml", variant), fmt.Sprintf("%s.yaml", variant)}
	var data []byte
	var err error

	for _, n := range names {
		p := filepath.Join(c.dir, n)
		if b, e := os.ReadFile(p); e == nil {
			data = b
			err = nil
			break
		} else {
			err = e
		}
	}

	if data == nil {
		data, err = readEmbeddedMessage(names)
	}
	if err != nil || data == nil {
		if variant != VariantDefault {
			return c.LoadVariant(VariantDefault)
		}
		return nil, fmt.Errorf("failed to load messages for variant %q: %w", variant, err)
	}

	var messages MessageSet
	if err := yaml.Unmarshal(data, &messages); err != nil {
		return nil, fmt.Errorf("failed to parse messages for variant %q: %w", variant, err)
	}

	if len(messages.Primary) == 0 || len(messages.Secondary) == 0 || len(messages.Footnotes) == 0 {
		return nil, fmt.Errorf("invalid message set for variant %q", variant)
	}

	return &messages, nil
}

// readEmbeddedMessage tries to read from embedded assets/messages.
func readEmbeddedMessage(names []string) ([]byte, error) {
	var lastErr error
	for _, n := range names {
		if b, err := assets.Messages.ReadFile(fmt.Sprintf("messages/%s", n)); err == nil {
			return b, nil
		} else {
			lastErr = err
		}
	}
	return nil, lastErr
}

// RenderMessage replaces mustache variables in the template with escaped values from the vars map.
// It uses {{key}} syntax for variables.
func (c *Messages) RenderMessage(template string, vars map[string]string) string {
	result := template
	for key, value := range vars {
		safe := html.EscapeString(value)
		result = strings.ReplaceAll(result, fmt.Sprintf("{{%s}}", key), safe)
	}
	return result
}
