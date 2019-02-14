// +build !go1.11

package file

import (
	"os"
)

func (f *File) _type() string {
	return "unknown"
}
