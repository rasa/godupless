// +build plan9

package file

import (
	"fmt"
)

// VolumeName @todo
func (f *File) VolumeName() (volume string, err error) {
	return fmt.Sprintf("%x", f.volumeID), nil
}
