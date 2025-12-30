package pipeline

import (
	"errors"
	"image"
	"os"
	"strconv"
	"strings"

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

const (
	minQuality   = 35
	qualityStep  = 7
	maxAttempts  = 6
)

func saveWithSize(
	img image.Image,
	outPath string,
	ext string,
	startQuality int,
	maxBytes int64,
	mode formats.Mode,
) error {

	q := startQuality
	tmp := outPath + ".tmp"

	for i := 0; i < maxAttempts && q >= minQuality; i++ {
		err := formats.Save(img, tmp, ext, q, mode)
		if err != nil {
			return err
		}

		info, err := os.Stat(tmp)
		if err != nil {
			return err
		}

		if info.Size() <= maxBytes {
			return os.Rename(tmp, outPath)
		}

		q -= qualityStep
	}

	_ = os.Rename(tmp, outPath)
	return nil
}