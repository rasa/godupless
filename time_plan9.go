// +build plan9

package main

import (
	"os"
	"syscall"
	"time"
)

// Ctime @todo
func Ctime(fi os.FileInfo) time.Time {
	// not supported
	return time.Unix(0, 0)
}

// Atime @todo
func Atime(fi os.FileInfo) time.Time {
	return time.Unix(int64(fi.Sys().(*syscall.Dir).Atime), 0)
}
