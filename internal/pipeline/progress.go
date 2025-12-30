package pipeline

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// progress bar from an older project
func printProgress(done, total int, start time.Time) {
	width := 20
	pct := float64(done) / float64(total)
	filled := int(pct * float64(width))

	bar := strings.Repeat("#", filled) + strings.Repeat(" ", width-filled)
	elapsed := time.Since(start).Seconds()

	var eta float64
	if done > 0 {
		eta = (elapsed / float64(done)) * float64(total-done)
	}

	fmt.Printf(
		"\r[%s] %3.0f%% | %d/%d images | ETA %.1fs",
		bar,
		pct*100,
		done,
		total,
		math.Max(0, eta),
	)
}
