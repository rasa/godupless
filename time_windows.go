// +build windows

package main

import (
	"os"
	"syscall"
	"time"
)

// Ctime @todo
func Ctime(fi os.FileInfo) time.Time {
	return time.Unix(0, fi.Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds())
}

// Atime @todo
func Atime(fi os.FileInfo) time.Time {
	return time.Unix(0, fi.Sys().(*syscall.Win32FileAttributeData).LastAccessTime.Nanoseconds())
}
