// +build aix dragonfly linux openbsd solaris

package file

import (
	"os"
	"syscall"
)

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

	f.volumeID = uint64(fi.Sys().(*syscall.Stat_t).Dev) // uint32
	f.fileID = fi.Sys().(*syscall.Stat_t).Ino
	f.atime = timespecToTime(fi.Sys().(*syscall.Stat_t).Atim)
	f.ctime = timespecToTime(fi.Sys().(*syscall.Stat_t).Ctim)
	f.nlinks = uint64(fi.Sys().(*syscall.Stat_t).Nlink) // uint32 (int16 on aix)
	return nil
}
