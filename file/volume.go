// +build !windows
// +build !plan9

package file

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
	if fi == nil {
		fi, err = os.Lstat(path)
		if err != nil {
			return err
		}
	}
	// see https://groups.google.com/forum/#!topic/golang-nuts/mu8XMmRXMOk
	rdev := uint64(fi.Sys().(*syscall.Stat_t).Rdev)
	devMap[rdev] = path
	return nil
}

// VolumeName @todo
func (f *File) VolumeName() (volume string, err error) {
	volume = fmt.Sprintf("%x", f.volumeID)
	if len(devMap) == 0 {
		err := filepath.Walk("/dev", visit)
		if err != nil {
			return volume, err
		}
	}
	vol, ok := devMap[f.volumeID]
	if ok {
		return vol, nil
	}
	return volume, nil
}
