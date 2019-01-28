// +build plan9

package main

import (
	"fmt"
	"os"
	"syscall"
)

// VolumeName @todo
func VolumeName(fi os.FileInfo, path string) string {
	dev := uint64(fi.Sys().(*syscall.Dir).Dev)
	typ := uint64(fi.Sys().(*syscall.Dir).Type)
	dev &= typ << 32
	return fmt.Sprintf("%x", dev)
}
