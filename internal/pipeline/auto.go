package pipeline

import (
	"image"
	"math"
)

func auto(img image.Image, ext string) int {
	b := img.Bounds()
	pixels := b.Dx() * b.Dy()
	edge, color := edgeDensity(img), colorVariance(img)
	q := baseQuality(ext)

	if edge > 0.18 {
		q += 15
	}
	if color < 0.04 {
		q += 10
	}
	if edge < 0.08 && color > 0.10 {
		q -= 10
	}

	switch {
	case pixels > 16_000_000:
		q -= 15
	case pixels > 8_000_000:
		q -= 10
	case pixels > 3_000_000:
		q -= 5
	}

	return clamp(q, 35, 92)
}


func baseQuality(ext string) int {
	switch ext {
	case ".avif":
		return 60
	case ".jpg", ".jpeg":
		return 78
	case ".png":
		return 100
	case ".webp":
		return 75
	default:
		return 80
	}
}


func clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}


func edgeDensity(img image.Image) float64 {
	b := img.Bounds()
	var edges, total float64

	for y := b.Min.Y + 1; y < b.Max.Y-1; y += 2 {
		for x := b.Min.X + 1; x < b.Max.X-1; x += 2 {
			l1 := luminance(img.At(x-1, y).RGBA())
			l2 := luminance(img.At(x+1, y).RGBA())
			if math.Abs(l1 - l2) > 0.08 {
				edges++
			}
			total++
		}
	}

	if total == 0 {
		return 0
	}
	return edges / total
}

func colorVariance(img image.Image) float64 {
	ib := img.Bounds()
	var sum, sumSq, count float64

	for y := ib.Min.Y; y < ib.Max.Y; y += 3 {
		for x := ib.Min.X; x < ib.Max.X; x += 3 {
			l := luminance(img.At(x, y).RGBA())
			sum += l
			sumSq += l * l
			count++
		}
	}

	if count == 0 {
		return 0
	}

	mean := sum / count
	return (sumSq / count) - (mean * mean)
}

func luminance(r, g, b, _ uint32) float64 {
	rr := float64(r) / 65535
	gg := float64(g) / 65535
	bb := float64(b) / 65535

	return 0.2126*rr + 0.7152*gg + 0.0722*bb
}
