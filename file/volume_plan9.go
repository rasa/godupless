// +build plan9

package file

import (
	"fmt"
)

// VolumeName @todo
func (f *File) VolumeName() (volume string, err error) {
	return fmt.Sprintf("%x", f.volumeID), nil
}

// GetVolumes @todo
func GetVolumes() ([]string, error) {
	return []string{"/"}, nil
}

// IsVolume @todo
func IsVolume(path string) bool {
	return path == "/"
}
