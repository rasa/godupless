// +build windows

package main

import (
	"os"
	"path/filepath"
)

// VolumeName @todo
func VolumeName(fi os.FileInfo, path string) string {
	return filepath.VolumeName(path)
}
