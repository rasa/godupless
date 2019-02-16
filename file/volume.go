// +build !windows
// +build !plan9

package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/rasa/godupless/util"
	
	"github.com/docker/docker/pkg/mount"
)

var devMap map[uint64]string
var volumes []string

func init() {
	devMap = make(map[uint64]string)
	volumes = make([]string, 0)
	// @todo load mounted volumes into volumes array
	volumes = append(volumes, "/")
}

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

func filterFunc(*mount.Info) (skip, stop bool) {
	return false, false
}

// GetVolumes @todo
func GetVolumes() ([]string, error) {
	var err error
	infoslice, err := mount.GetMounts(filterFunc)
	util.Dump("infoslice=", infoslice)
	
	/*
		if len(volumeMap) == 0 {
			err := loadVolumeMap()
			if err != nil {
				return volumes, err
			}
		}
	*/
	return volumes, err
}
