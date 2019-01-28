// +build windows

package main

import (
	"os"
)

// GetDev @todo
func GetDev(fi os.FileInfo, path string) string {
	return substring(path, 0, 2)
}
