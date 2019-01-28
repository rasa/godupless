// +build dragonfly linux openbsd solaris

package main

import (
	"os"
	"syscall"
	"time"
)

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

// Ctime @todo
func Ctime(fi os.FileInfo) time.Time {
	return timespecToTime(fi.Sys().(*syscall.Stat_t).Ctim)
}

// Atime @todo
func Atime(fi os.FileInfo) time.Time {
	return timespecToTime(fi.Sys().(*syscall.Stat_t).Atim)
}
