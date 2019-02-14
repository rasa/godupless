package file

import (
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"os"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

// File @todo
type File struct {
	/*
		// VolumeID @todo
		VolumeID string `csv:"volumeID"`
		// FileID @todo
		FileID string `csv:"fileID"`
		// Path @todo
		Path string `csv:"path"`
		// Size @todo
		Size uint64 `csv:"size"`
		// Modified @todo
		Modified time.Time `csv:"modified"`
		// Hash @todo
		Hash string `csv:"hash"`
		// Mode @todo
		Mode string `csv:"mode,omitempty"`
		// Created @todo
		Created string `csv:"created,omitempty"`
		// Accessed @todo
		Accessed string `csv:"accessed,omitempty"`
		// Valid @todo
		Valid bool `csv:"-"`
	*/
	path  string
	size  uint64
	mode  os.FileMode
	mtime time.Time

	volumeID uint64
	fileID   uint64
	atime    time.Time
	ctime    time.Time
	nlinks   uint64
	// Sys() interface{}

	h   hash.Hash
	fh  *os.File
	pos uint64
	eof bool
	err error
}

// NewFile @todo
func NewFile(path string, fi os.FileInfo, h hash.Hash) (f *File, err error) {
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}
	f = &File{path: path, h: h}
	err = f.stat(fi)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// Name @todo
func (f *File) Name() string {
	return basename(f.path)
}

// Path @todo
func (f *File) Path() string {
	return f.path
}

// Size @todo
func (f *File) Size() uint64 {
	return f.size
}

// Mode @todo
func (f *File) Mode() os.FileMode {
	return f.mode
}

// ModTime @todo
func (f *File) ModTime() time.Time {
	return f.mtime
}

// IsDir @todo
func (f *File) IsDir() bool {
	return f.mode&os.ModeDir != 0
}

// IsSymlink @todo
func (f *File) IsSymlink() bool {
	return f.mode&os.ModeSymlink != 0
}

// IsRegular @todo
func (f *File) IsRegular() bool {
	return f.mode&os.ModeType == 0
}

// VolumeID @todo
func (f *File) VolumeID() uint64 {
	return f.volumeID
}

// FileID @todo
func (f *File) FileID() uint64 {
	return f.fileID
}

// UniqueID @todo
func (f *File) UniqueID() string {
	return fmt.Sprintf("%016x%016x", f.volumeID, f.fileID)
}

// Atime @todo
func (f *File) Atime() time.Time {
	return f.atime
}

// Ctime @todo
func (f *File) Ctime() time.Time {
	return f.ctime
}

// Type @todo
func (f *File) Type() string {
	if f.mode&os.ModeType == 0 {
		return "regular file"
	}
	if f.mode&os.ModeDir != 0 {
		return "directory"
	}
	if f.mode&os.ModeSymlink != 0 {
		return "Symlink"
	}
	if f.mode&os.ModeDevice != 0 {
		return "device"
	}
	if f.mode&os.ModeNamedPipe != 0 {
		return "named pipe"
	}
	if f.mode&os.ModeSocket != 0 {
		return "socket"
	}
	return f._type()
}

// Reset @todo
func (f *File) Reset() {
	f.Close()
	f.pos = 0
	f.err = nil
	f.eof = false
	f.h.Reset()
}

// Open @todo
func (f *File) Open() error {
	f.Reset()
	return f.Reopen()
}

// Reopen @todo
func (f *File) Reopen() error {
	if f.err != nil {
		// don't reopen if there's been an error
		return f.err
	}
	if f.fh != nil {
		// no need to reopen if it's already open
		return nil
	}
	f.fh, f.err = os.Open(f.path)
	if f.err != nil {
		return f.err
	}
	if f.pos == 0 {
		// we're already at the beginning of the file
		return nil
	}
	fmt.Printf("Seeking to %d\n", f.pos)
	pos, err := f.fh.Seek(int64(f.pos), os.SEEK_SET)
	if err != nil {
		if err == io.EOF {
			f.eof = true
			return io.EOF
		}
		f.err = err
		return err
	}
	if uint64(pos) != f.pos {
		return fmt.Errorf("seek failed: expected %d, got %d", f.pos, pos)
	}
	return nil
}

// Read @todo
func (f *File) Read(bytes uint64) (err error) {
	if f.eof {
		return io.EOF
	}
	if f.err != nil {
		return f.err
	}
	if f.fh == nil {
		fmt.Printf("file.Read(): Reopening %s\n", f.path)
		err = f.Reopen()
		if err != nil {
			return err
		}
	}
	written, err := io.CopyN(f.h, f.fh, int64(bytes))
	f.pos += uint64(written)
	if err != nil {
		if err == io.EOF {
			f.eof = true
		} else {
			f.err = err
		}
	}
	return err
}

// EOF @todo
func (f *File) EOF() bool {
	return f.eof
}

// Err @todo
func (f *File) Err() error {
	return f.err
}

// Sum @todo
func (f *File) Sum() []byte {
	if f.h == nil {
		return nil
	}
	return f.h.Sum(nil)
}

// Hash @todo
func (f *File) Hash() string {
	if f.h == nil {
		return ""
	}
	return fmt.Sprintf("%x", f.h.Sum(nil))
}

// Pos @todo
func (f *File) Pos() uint64 {
	return f.pos
}

// Opened @todo
func (f *File) Opened() bool {
	return f.fh != nil
}

// Close @todo
func (f *File) Close() (err error) {
	if f.fh != nil {
		err = f.fh.Close()
		f.fh = nil
	}
	return err
}

// debug functions:

// String @todo
func (f *File) String() string {
	return fmt.Sprintf("path: %s, size: %d, mode: %s, ctime: %s, mtime: %s, atime: %s, volumeID: %016x, fileID: %016x, nlinks: %d, pos: %d, eof: %v, hash: %s",
		f.path,
		f.size,
		f.mode.String(),
		f.ctime.Format(time.RFC3339), // RFC3339Nano
		f.mtime.Format(time.RFC3339),
		f.atime.Format(time.RFC3339),
		f.volumeID,
		f.fileID,
		f.nlinks,
		f.pos,
		f.eof,
		f.Hash())
}

// MarshalJSON @todo
func (f *File) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}
