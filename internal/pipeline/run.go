package pipeline

import (
	"os"

	"github.com/kaustubha-chaturvedi/yst-img/internal/formats"
)

func Run(input, output string, quality, workers int, format string, maxSize string, mode formats.Mode) error {
	info, err := os.Stat(input)
	if err != nil {
		return err
	}

	if info.IsDir() {
		// to impl
		// i am planning bathc ops too
	}
	// to impl
	return nil
}
