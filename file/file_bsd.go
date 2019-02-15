// +build darwin freebsd netbsd

package file

import (
	"os"
	"syscall"

	"github.com/rasa/godupless/util"
)

// These stat field names may need to be renamed later
// see https://github.com/golang/go/issues/29393

func (f *File) stat(fi os.FileInfo) (err error) {
	if fi == nil {
		fi, err = os.Lstat(f.path)
		if err != nil {
			return err
		}
	}
	f.size = uint64(fi.Size())
	f.mode = fi.Mode()
	f.mtime = fi.ModTime()

	s, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return errors.New("conversion to *syscall.Stat_t failed")
	}
	f.volumeID = uint64(s.Dev) // int32
	f.fileID = uint64(s.Ino)   // uint32
	f.atime = util.TimespecToTime(s.Atimespec)
	f.ctime = util.TimespecToTime(s.Ctimespec)
	f.nlinks = uint64(s.Nlink) // uint16
	return nil
}
