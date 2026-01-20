//go:build avif
// +build avif

package formats

/*
#cgo linux pkg-config: libavif
#cgo darwin pkg-config: libavif
#cgo windows LDFLAGS: -L${SRCDIR}/../embed -lavif

#include <stdlib.h>
#include <stdint.h>
#include <avif/avif.h>
*/
import "C"

import (
	"errors"
	"image"
	"os"
	"unsafe"

	_ "github.com/kaustubha-chaturvedi/yst-img/internal/native"
)

func saveAVIF(img image.Image, path string, quality int) error {
	b := img.Bounds()
	w := b.Dx()
	h := b.Dy()

	encoder := C.avifEncoderCreate()
	if encoder == nil {
		return errors.New("avifEncoderCreate failed")
	}
	defer C.avifEncoderDestroy(encoder)

	encoder.maxThreads = 4
	encoder.quality = avifQuality(quality)


	avifImg := C.avifImageCreate(
		C.uint32_t(w),
		C.uint32_t(h),
		8,
		C.AVIF_PIXEL_FORMAT_YUV420,
	)
	if avifImg == nil {
		return errors.New("avifImageCreate failed")
	}
	defer C.avifImageDestroy(avifImg)

	if C.avifImageAllocatePlanes(avifImg, C.AVIF_PLANES_YUV) != C.AVIF_RESULT_OK {
		return errors.New("avifImageAllocatePlanes failed")
	}

	var rgb C.avifRGBImage
	C.avifRGBImageSetDefaults(&rgb, avifImg)
	rgb.format = C.AVIF_RGB_FORMAT_RGBA
	C.avifRGBImageAllocatePixels(&rgb)
	defer C.avifRGBImageFreePixels(&rgb)

	pixels := (*[1 << 30]uint8)(unsafe.Pointer(rgb.pixels))[:w*h*4]
	i := 0
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, bb, a := img.At(x, y).RGBA()
			pixels[i+0] = uint8(r >> 8)
			pixels[i+1] = uint8(g >> 8)
			pixels[i+2] = uint8(bb >> 8)
			pixels[i+3] = uint8(a >> 8)
			i += 4
		}
	}

	if C.avifImageRGBToYUV(avifImg, &rgb) != C.AVIF_RESULT_OK {
		return errors.New("avifImageRGBToYUV failed")
	}

	var out C.avifRWData
	C.avifRWDataRealloc(&out, 0)
	defer C.avifRWDataFree(&out)

	if C.avifEncoderWrite(encoder, avifImg, &out) != C.AVIF_RESULT_OK {
		return errors.New("avifEncoderWrite failed")
	}

	return os.WriteFile(
		path,
		C.GoBytes(unsafe.Pointer(out.data), C.int(out.size)),
		0644,
	)
}

func avifQuality(q int) C.int {
	if q < 1 {
		q = 1
	}
	if q > 100 {
		q = 100
	}
	return C.int(63 - (q * 63 / 100))
}
