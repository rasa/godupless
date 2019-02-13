// +build darwin freebsd netbsd

package file

import (
	"os"
	"syscall"
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

	f.volumeID = uint64(fi.Sys().(*syscall.Stat_t).Dev) // int32
	f.fileID = uint64(fi.Sys().(*syscall.Stat_t).Ino)   // uint32
	f.atime = timespecToTime(fi.Sys().(*syscall.Stat_t).Atimespec)
	f.ctime = timespecToTime(fi.Sys().(*syscall.Stat_t).Ctimespec)
	f.nlinks = uint64(fi.Sys().(*syscall.Stat_t).Nlink) // uint16
	return nil
}
