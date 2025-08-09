package og

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	MaxTextWidth  = 40
	FileExtension = ".png"
	HashLength    = 12
	MaxFilenameID = 10000
)

func WrapText(text string, maxWidth int) string {
	if len(text) <= maxWidth {
		return text
	}

	words := strings.Fields(text)
	var lines []string
	var currentLine []string
	currentLength := 0

	for _, word := range words {
		if currentLength+len(word)+1 > maxWidth {
			if len(currentLine) > 0 {
				lines = append(lines, strings.Join(currentLine, " "))
				currentLine = []string{word}
				currentLength = len(word)
			} else {
				lines = append(lines, word)
				currentLine = []string{}
				currentLength = 0
			}
		} else {
			currentLine = append(currentLine, word)
			currentLength += len(word) + 1
		}
	}

	if len(currentLine) > 0 {
		lines = append(lines, strings.Join(currentLine, " "))
	}

	return strings.Join(lines, "\n")
}

func GenerateCacheKey(values ...string) string {
	content := strings.Join(values, ":")
	hash := fmt.Sprintf("%x", md5.Sum([]byte(content)))
	return hash[:HashLength]
}

func CheckCache(cacheKey, publicPath string) (string, bool) {
	filename := fmt.Sprintf("%s%s", cacheKey, FileExtension)
	fullPath := filepath.Join(publicPath, filename)

	if _, err := os.Stat(fullPath); err == nil {
		return fullPath, true
	}

	return "", false
}

func GenerateSarcasticFilename(count, target string) string {
	sarcasticWords := []string{
		"corporate-silence", "professional-void", "executive-quiet",
		"strategic-ignore", "enterprise-mute", "business-shush",
		"silence-metrics", "quiet-kpis", "mute-analytics",
		"wisdom-declined", "insights-rejected", "thoughts-ignored",
	}

	hash := 0
	for _, char := range count {
		hash = hash*31 + int(char)
	}
	if hash < 0 {
		hash = -hash
	}

	word := sarcasticWords[hash%len(sarcasticWords)]

	if target != "" {
		return fmt.Sprintf("%s-%s-ignored%s", word, strings.ToLower(target), FileExtension)
	}

	return fmt.Sprintf("%s-%d%s", word, hash%MaxFilenameID, FileExtension)
}
