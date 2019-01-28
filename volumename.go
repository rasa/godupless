// +build android darwin dragonfly freebsd js,wasm linux nacl netbsd openbsd solaris

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

var devMap map[uint64]string

func init() {
	devMap = make(map[uint64]string)
}

func visit(path string, fi os.FileInfo, err error) error {
	rdev := uint64(fi.Sys().(*syscall.Stat_t).Rdev)
	devMap[rdev] = path
	return nil
}

// GetDev @todo
func VolumeName(fi os.FileInfo, path string) string {
	if len(devMap) == 0 {
		err := filepath.Walk("/dev", visit)
		if err != nil {
			panic(err)
		}
	}
	dev := uint64(fi.Sys().(*syscall.Stat_t).Dev)
	name, ok := devMap[dev]
	if ok {
		return name
	}
	return fmt.Sprintf("%x", dev)
}
