package pipeline

import (
	"path/filepath"
	"strings"

	"github.com/kaustubha-chaturvedi/yst-img/internal/formats"
)

type OutputSpec struct {
	Path string
	Ext  string
}


func resolveOutput(
	output string,
	maxSizeBytes int64,
	mode formats.Mode,
) OutputSpec {

	ext := strings.ToLower(filepath.Ext(output))
	if mode == formats.Compress  && ext == ".png"{
		newExt := formats.FallbackFormatForPngCompress(ext)
		if newExt != ext {
			output = strings.TrimSuffix(output, ext) + newExt
			ext = newExt
		}
	}

	if maxSizeBytes > 0 {
		newExt := formats.FallbackFormatForPngCompress(ext)
		if newExt != ext {
			output = strings.TrimSuffix(output, ext) + newExt
			ext = newExt
		}
	}
	
	return OutputSpec{
		Path: output,
		Ext:  ext,
	}
}
