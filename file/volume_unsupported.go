// +build aix js nacl

package file

import (
	"fmt"
	"runtime"
)

// VolumeName @todo
func (f *File) VolumeName() (string, error) {
	return fmt.Sprintf("%x", f.volumeID), nil
}

// GetVolumes @todo
func GetVolumes() ([]string, error) {
	fmt.Printf("GetVolumes() not supported on %s/%s\n", runtime.GOOS, runtime.GOARCH)
	return []string{}, nil
}

// IsVolume @todo
func IsVolume(path string) bool {
	return path == "/"
}
