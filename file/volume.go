// +build !aix
// +build !js
// +build !nacl
// +build !plan9
// +build !windows

// aix,nacl
//   docker/docker/vendor/github.com/sirupsen/logrus/terminal_check_notappengine.go:15:10: undefined: terminal.IsTerminal
// js:
// 	 docker/docker/pkg/mount/unmount_unix.go:5:8: build constraints exclude all Go files in d:/go/src/github.com/docker/docker/vendor/golang.org/x/sys/unix

package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/rasa/godupless/util"

	"github.com/docker/docker/pkg/mount"
)

var devMap map[uint64]string
var volumes []string
var volMap map[string]*mount.Info

func init() {
	devMap = make(map[uint64]string)
	volumes = make([]string, 0)
	volMap = make(map[string]*mount.Info)
}

var donce sync.Once
var devErr error

func visit(path string, fi os.FileInfo, err error) error {
	if fi == nil {
		fi, err = os.Lstat(path)
		if err != nil {
			return err
		}
	}
	// see https://groups.google.com/forum/#!topic/golang-nuts/mu8XMmRXMOk
	s, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return errors.New("conversion to *syscall.Stat_t failed")
	}

	rdev := uint64(s.Rdev)
	devMap[rdev] = path
	return nil
}

func loadDevMap() {
	devErr = filepath.Walk("/dev", visit)
	util.Dump("devMap=", devMap)
}

// VolumeName @todo
func (f *File) VolumeName() (volume string, err error) {
	volume = fmt.Sprintf("%x", f.volumeID)
	donce.Do(loadDevMap)
	if devErr != nil {
		return volume, devErr
	}
	vol, ok := devMap[f.volumeID]
	if ok {
		return vol, nil
	}
	return volume, nil
}

func filterFunc(i *mount.Info) (skip, stop bool) {
	return false, false
}

var pathToIgnore = []string{
	"/dev",
	"/proc",
	"/run",
	"/sys",
}

var vonce sync.Once
var volErr error

func loadVolumeMap() {
	var mounts []*mount.Info
	mounts, volErr = mount.GetMounts(filterFunc)
	util.Dump("mounts=", mounts)
Loop:
	for _, mount := range mounts {
		for _, path := range pathToIgnore {
			if strings.HasPrefix(mount.Mountpoint, path) {
				continue Loop
			}
		}
		volMap[mount.Mountpoint] = mount
		volumes = append(volumes, mount.Mountpoint)
	}
}

// GetVolumes @todo
func GetVolumes() ([]string, error) {
	vonce.Do(loadVolumeMap)
	util.Dump("volumes=", volumes)
	return volumes, volErr
}

// IsVolume @todo
func IsVolume(path string) bool {
	vonce.Do(loadVolumeMap)
	_, ok := volMap[path]
	return ok
}
