// +build !go1.11

package file

func (f *File) _type() string {
	return "unknown"
}
