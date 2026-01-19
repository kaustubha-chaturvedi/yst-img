//go:build windows

package native

import (
	"os"
	"path/filepath"

	"github.com/kaustubha-chaturvedi/yst-img/internal/embed"
)

func init() {
	dir := filepath.Join(os.TempDir(), "yst-img", "dlls")
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic("failed to create temp dir for DLLs: " + err.Error())
	}

	entries, err := embed.DLLs.ReadDir(".")
	if err != nil {
		panic("failed to read embedded DLLs: " + err.Error())
	}

	for _, e := range entries {
		if filepath.Ext(e.Name()) != ".dll" {
			continue
		}
		data, err := embed.DLLs.ReadFile(e.Name())
		if err != nil {
			panic("failed to read DLL " + e.Name() + ": " + err.Error())
		}
		if err := os.WriteFile(filepath.Join(dir, e.Name()), data, 0644); err != nil {
			panic("failed to write DLL " + e.Name() + ": " + err.Error())
		}
	}

	path := os.Getenv("PATH")
	if err := os.Setenv("PATH", dir+";"+path); err != nil {
		panic("failed to update PATH: " + err.Error())
	}
}