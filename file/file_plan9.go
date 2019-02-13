// +build plan9

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

	f.volumeID = uint64(fi.Sys().(*syscall.Dir).Type)<<32 + uint64(fi.Sys().(*syscall.Dir).Dev)
	f.fileID = fi.Sys().(*syscall.Dir).Qid.Path
	f.atime = time.Unix(int64(fi.Sys().(*syscall.Dir).Atime), 0)
	// not supported:
	// f.ctime = time.Unix(0, 0)
	// f.nlinks = 0
	return nil
}
