// +build plan9

package file

import (
	"os"
	"syscall"
	"time"

	"github.com/cespare/xxhash"
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

	s, ok := fi.Sys().(*syscall.Dir)
	if !ok {
		return errDirConversion
	}
	f.volumeID = uint64(s.Type)<<32 + uint64(s.Dev)
	f.fileID = s.Qid.Path
	f.atime = time.Unix(int64(s.Atime), 0)
	// not supported:
	// f.ctime = time.Unix(0, 0)
	// f.nlinks = 0
	f.uid = xxhash.Sum64([]byte(s.Uid))
	f.gid = xxhash.Sum64([]byte(s.Gid))
	return nil
}
