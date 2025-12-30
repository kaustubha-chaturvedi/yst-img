package formats

import (
	"fmt"
	"image"
)

type Mode int
const (
	Convert Mode = iota
	Compress
	Resize
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
