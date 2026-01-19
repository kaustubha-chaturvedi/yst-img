//go:build windows
package embed

import "embed"

//go:embed *.dll
var DLLs embed.FS
