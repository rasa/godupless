// +build darwin,freebsd,netbsd,openbsd

package main

import (
	"os"
	"syscall"
	"time"
)

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

// These stat field names may need to be renamed later
// see https://github.com/golang/go/issues/29393

// Ctime @todo
func Ctime(fi os.FileInfo) time.Time {
	return timespecToTime(fi.Sys().(*syscall.Stat_t).Ctim)
}

// Atime @todo
func Atime(fi os.FileInfo) time.Time {
	return timespecToTime(fi.Sys().(*syscall.Stat_t).Atim)
}
