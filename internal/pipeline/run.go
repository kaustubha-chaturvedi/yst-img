package pipeline

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/kaustubha-chaturvedi/yst-img/internal/formats"
)


func parseSize(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}

	s = strings.ToLower(strings.TrimSpace(s))
	mult := int64(1)

	switch {
	case strings.HasSuffix(s, "k"):
		mult = 1024
		s = strings.TrimSuffix(s, "k")
	case strings.HasSuffix(s, "m"):
		mult = 1024 * 1024
		s = strings.TrimSuffix(s, "m")
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, errors.New("invalid --max-size value")
	}

	return v * mult, nil
}


func Run(input, output string, quality, workers int, format string, maxSize string, mode formats.Mode) error {
	info, err := os.Stat(input)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return runBatch(input, output, quality, workers, format, maxSize, mode)
	}

	return runSingle(input, output, quality, maxSize, mode)
}


func Convert(input, output string, quality int, mode formats.Mode) error {
	img, err := imaging.Open(input)
	if err != nil {
		return err
	}
	ext := strings.ToLower(filepath.Ext(output))
	return formats.Save(img, output, ext, quality, mode)
}


func Resize(input, output string, width int) error {
	img, err := imaging.Open(input)
	if err != nil {
		return fmt.Errorf("open failed: %w", err)
	}
	resized := imaging.Resize(img, width, 0, imaging.Lanczos)
	ext := strings.ToLower(filepath.Ext(output))
	return formats.Save(resized, output, ext, 100, formats.Resize)
}