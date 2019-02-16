// +build aix dragonfly linux openbsd solaris

package file

import (
	"errors"
	"os"
	"syscall"

	"github.com/rasa/godupless/util"
)

// MinDirLength @todo
const MinDirLength = 2

// ExcludePaths @todo
var ExcludePaths = []string{
	"^/dev$",
	"^/proc$",
	"^/run$",
	"^/sys$",
}

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

	f.volumeID = uint64(s.Dev) // uint32
	f.fileID = s.Ino
	f.atime = util.TimespecToTime(s.Atim)
	f.ctime = util.TimespecToTime(s.Ctim)
	f.nlinks = uint64(s.Nlink) // uint32 (int16 on aix)
	return nil
}
