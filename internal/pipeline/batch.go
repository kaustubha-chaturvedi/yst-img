package pipeline

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/disintegration/imaging"
	"github.com/kaustubha-chaturvedi/yst-img/internal/formats"
)

func runBatch(src, dst string, quality, workers int, format string, maxSize string, mode formats.Mode) error {
	files := collectImages(src)

	fmt.Printf("[batch] found %d images in %s\n", len(files), src)

	_ = os.MkdirAll(dst, 0755)

	start := time.Now()
	var done int64

	jobs := make(chan string)
	wg := sync.WaitGroup{}

	for range workers {
		wg.Go(func() {
			for file := range jobs {
				processBatchFile(file, dst, quality, format, maxSize, mode)
				atomic.AddInt64(&done, 1)
				printProgress(int(done), len(files), start)
			}
		})
	}

	for _, f := range files {
		jobs <- f
	}
	close(jobs)
	wg.Wait()

	fmt.Println("\n[batch] done")
	return nil
}

func collectImages(root string) []string {
	var files []string
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".webp" {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func processBatchFile(path, dst string, quality int, format string, maxSize string, mode formats.Mode) {
	img, err := imaging.Open(path)
	if err != nil {
		return
	}

	ext := strings.ToLower(filepath.Ext(path))
	if format == "" {
		format = strings.TrimPrefix(ext, ".")
	}
	name := strings.TrimSuffix(filepath.Base(path), ext)
	
	targetBytes, _ := parseSize(maxSize)
	
	spec := resolveOutput(
		filepath.Join(dst, name+"."+format),
		targetBytes,
		mode,
	)

	if quality == 0 {
		// i am planning to make auto mode for image compression
		// quality = auto(img, spec.Ext)
	}
		
	if targetBytes > 0 {
		_ = formats.SaveWithSize(img, spec.Path, spec.Ext, quality, targetBytes, mode)
		return
	}

	_ = formats.Save(img, spec.Path, "."+format, quality, mode)
}
