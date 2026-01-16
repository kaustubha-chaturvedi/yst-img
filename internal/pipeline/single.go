package pipeline

import (
	"github.com/disintegration/imaging"
	"github.com/kaustubha-chaturvedi/yst-img/internal/formats"
)

func runSingle(input, output string, quality int, maxSize string, mode formats.Mode) error {
	img, err := imaging.Open(input)

	if err != nil {
		return err
	}

	
	
	targetBytes, err := parseSize(maxSize)
	if err != nil {
		return err
	}
	
	spec := resolveOutput(output, targetBytes, mode)

	if quality == 0 {
		quality = auto(img, spec.Ext)
	}
	
	if targetBytes > 0 {
		return formats.SaveWithSize(img, spec.Path, spec.Ext, quality, targetBytes, mode)
	}

	return formats.Save(img, spec.Path, spec.Ext, quality, mode)
}

