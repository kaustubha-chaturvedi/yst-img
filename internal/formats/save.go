package formats

import (
	"fmt"
	"image"
	"os"
)


type Mode int
const (
	Convert Mode = iota
	Compress
	Resize
)

const (
	minQuality   = 35
	qualityStep  = 7
	maxAttempts  = 6
)

func FallbackFormatForPngCompress(ext string) string {
	switch ext {
	case ".png":
		return ".jpg"
	default:
		return ext
	}
}


func Save(img image.Image, path, ext string, quality int, mode Mode) error {
	if mode == Compress && ext ==".png" {
		return saveJPEG(img, path, quality)
	}
	switch ext {
	case ".jpg", ".jpeg":
		return saveJPEG(img, path, quality)
	case ".png":
		return savePNG(img, path)
	case ".webp":
		return saveWebP(img, path, quality)
	default:
		return fmt.Errorf("unsupported format: %s", ext)
	}
}


func SaveWithSize(
	img image.Image,
	outPath string,
	ext string,
	startQuality int,
	maxBytes int64,
	mode Mode,
) error {

	q := startQuality
	tmp := outPath + ".tmp"

	for i := 0; i < maxAttempts && q >= minQuality; i++ {
		err := Save(img, tmp, ext, q, mode)
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