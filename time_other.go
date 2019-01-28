// +build nacl js,wasm 

package main

import (
	"os"
	"syscall"
	"time"
)

// Ctime @todo
func Ctime(fi os.FileInfo) time.Time {
	secs := fi.Sys().(*syscall.Stat_t).Ctime
	nsecs := fi.Sys().(*syscall.Stat_t).CtimeNsec
	return time.Unix(secs, nsecs)
}

// Atime @todo
func Atime(fi os.FileInfo) time.Time {
	secs := fi.Sys().(*syscall.Stat_t).Atime
	nsecs := fi.Sys().(*syscall.Stat_t).AtimeNsec
	return time.Unix(secs, nsecs)
}
