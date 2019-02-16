// +build windows

package file

import (
	"errors"
	"os"
	"sync"
	"syscall"
	"time"
)

// MinDirLength @todo
const MinDirLength = 3

// ExcludePaths @todo
var ExcludePaths = []string{
	`(?i)^[A-Z]:/$Recycle\.bin`,
	`(?i)^[A-Z]:/System Volume Information`,
}

var mutex sync.Mutex

func (f *File) stat(fi os.FileInfo) (err error) {
	mutex.Lock()
	defer mutex.Unlock()

	if fi == nil {
		fi, err = os.Lstat(f.path)
		if err != nil {
			return err
		}
	}

	f.size = uint64(fi.Size())
	f.mode = fi.Mode()
	f.mtime = fi.ModTime()

	s, ok := fi.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return errors.New("conversion to *syscall.Win32FileAttributeData failed")
	}
	f.atime = time.Unix(0, s.LastAccessTime.Nanoseconds())
	f.ctime = time.Unix(0, s.CreationTime.Nanoseconds())

	pathp, err := syscall.UTF16PtrFromString(f.path)
	if err != nil {
		return err
	}
	h, err := syscall.CreateFile(pathp, 0, 0, nil, syscall.OPEN_EXISTING, syscall.FILE_FLAG_BACKUP_SEMANTICS, 0)
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(h)
	var i syscall.ByHandleFileInformation
	err = syscall.GetFileInformationByHandle(h, &i)
	if err != nil {
		return err
	}

	// f.size = uint64(i.FileSizeHigh << 32) + uint64(i.FileSizeLow)
	f.volumeID = uint64(i.VolumeSerialNumber)
	f.fileID = (uint64(i.FileIndexHigh) << 32) + uint64(i.FileIndexLow)
	f.nlinks = uint64(i.NumberOfLinks)
	return nil
}

// basename removes trailing slashes and the leading
// directory name and drive letter from path name.
func basename(name string) string {
	// Remove drive letter
	if len(name) == 2 && name[1] == ':' {
		name = "."
	} else if len(name) > 2 && name[1] == ':' {
		name = name[2:]
	}
	i := len(name) - 1
	// Remove trailing slashes
	for ; i > 0 && (name[i] == '/' || name[i] == '\\'); i-- {
		name = name[:i]
	}
	// Remove leading directory name
	for i--; i >= 0; i-- {
		if name[i] == '/' || name[i] == '\\' {
			name = name[i+1:]
			break
		}
	}
	return name
}
