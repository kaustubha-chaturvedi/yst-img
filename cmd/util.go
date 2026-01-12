package cmd

import (
	"os"
	"path/filepath"
	"strings"
)

func defaultOutput(input string, def string) string {
	info, err := os.Stat(input)
	if err != nil {
		return input + def
	}
	
	if info.IsDir() {
		return input + def
	}

	ext := filepath.Ext(input)
	base := strings.TrimSuffix(input, ext)

	if ext == "" {
		return base + def
	}

	return base + def + ext
}