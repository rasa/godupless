// +build go1.11

package file

import (
	"os"
)

func (f *File) _type() string {
	// added in go 1.11:
	if f.mode&os.ModeIrregular != 0 {
		return "irregular"
	}
	return "unknown"
}
