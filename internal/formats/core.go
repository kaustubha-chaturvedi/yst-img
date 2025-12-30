package formats

import (
	"image"
	"image/png"
	"image/jpeg"
	"os"
	"github.com/chai2010/webp"
)


func savePNG(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}


func saveJPEG(img image.Image, path string, q int) error {
	if q < 1 {
		q = 1
	}
	if q > 100 {
		q = 100
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return jpeg.Encode(f, img, &jpeg.Options{Quality: q})
}


func saveWebP(img image.Image, path string, quality int) error {
	if quality <= 0 || quality > 100 {
		quality = 75
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	opts := &webp.Options{
		Lossless: false,
		Quality:  float32(quality),
	}

	return webp.Encode(f, img, opts)
}