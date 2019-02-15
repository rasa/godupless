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

	s, ok := fi.Sys().(*syscall.Dir)
	if !ok {
		return nil, errors.New("conversion to *syscall.Dir failed")
	}
	f.volumeID = uint64(s.Type)<<32 + uint64(s.Dev)
	f.fileID = s.Qid.Path
	f.atime = time.Unix(int64(s.Atime), 0)
	// not supported:
	// f.ctime = time.Unix(0, 0)
	// f.nlinks = 0
	return nil
}
