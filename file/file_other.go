// +build nacl js,wasm

package file

import (
	"os"
	"syscall"
	"time"
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

	f.volumeID = uint64(fi.Sys().(*syscall.Stat_t).Dev) // int64
	f.fileID = fi.Sys().(*syscall.Stat_t).Ino
	f.atime = time.Unix(fi.Sys().(*syscall.Stat_t).Atime, fi.Sys().(*syscall.Stat_t).AtimeNsec)
	f.ctime = time.Unix(fi.Sys().(*syscall.Stat_t).Ctime, fi.Sys().(*syscall.Stat_t).CtimeNsec)
	f.nlinks = uint64(fi.Sys().(*syscall.Stat_t).Nlink) // uint32
	return nil
}
