package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kaustubha-chaturvedi/yst-img/cmd"
)


var embeddedMetadata = `{
	"name": "image tools",
	"domain": "image",
	"alias": "img",
	"version": "1.0.0",
	"formats": ["png", "jpeg", "webp", "avif"],
	"features": [
	  "convert",
	  "compress",
	  "auto-quality",
	  "batch",
	  "parallel",
	  "avif-av1"
	]
  }`

func main() {
	if len(os.Args) > 1 && os.Args[1] == "__yst_metadata" {
		meta := getMetadata()
		json.NewEncoder(os.Stdout).Encode(meta)
		return
	}

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func getMetadata() map[string]interface{} {
	if embeddedMetadata != "" {
		var meta map[string]interface{}
		if err := json.Unmarshal([]byte(embeddedMetadata), &meta); err == nil {
			return meta
		}
	}
	return map[string]interface{}{
		"name":     "image tools",
		"domain":   "image",
		"alias":    "img",
		"version":  "0.1.0",
		"commands": []map[string]interface{}{},
	}
}