// +build android darwin dragonfly freebsd js,wasm linux nacl netbsd openbsd solaris

package main

import (
	"fmt"
	"os"
	"syscall"
)

// GetDev @todo
func GetDev(fi os.FileInfo, path string) string {
	return fmt.Sprintf("%x", fi.Sys().(*syscall.Stat_t).Dev)
}
