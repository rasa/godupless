// +build nacl js,wasm

package file

import (
	"os"
	"syscall"
	"time"
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
		return errStatConversion
	}

	f.volumeID = uint64(s.Dev) // int64
	f.fileID = s.Ino
	f.atime = time.Unix(s.Atime, s.AtimeNsec)
	f.ctime = time.Unix(s.Ctime, s.CtimeNsec)
	f.nlinks = uint64(s.Nlink) // uint32
	f.uid = uint64(s.Uid)
	f.gid = uint64(s.Gid)
	return nil
}
