package file

import (
	"hash"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestFil(t *testing.T) {
	t.Logf("TestFile")
}

func TestNewFile(t *testing.T) {
	type args struct {
		path string
		fi   os.FileInfo
		h    hash.Hash
	}
	tests := []struct {
		name    string
		args    args
		wantF   *File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF, err := NewFile(tt.args.path, tt.args.fi, tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotF, tt.wantF) {
				t.Errorf("NewFile() = %v, want %v", gotF, tt.wantF)
			}
		})
	}
}

func TestFile_Name(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Name(); got != tt.want {
				t.Errorf("File.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Dir(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Dir(); got != tt.want {
				t.Errorf("File.Dir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Path(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Path(); got != tt.want {
				t.Errorf("File.Path() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Size(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Size(); got != tt.want {
				t.Errorf("File.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Mode(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   os.FileMode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Mode(); got != tt.want {
				t.Errorf("File.Mode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_ModTime(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.ModTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.ModTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_IsDir(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.IsDir(); got != tt.want {
				t.Errorf("File.IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_IsSymlink(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.IsSymlink(); got != tt.want {
				t.Errorf("File.IsSymlink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_IsRegular(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.IsRegular(); got != tt.want {
				t.Errorf("File.IsRegular() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_VolumeID(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.VolumeID(); got != tt.want {
				t.Errorf("File.VolumeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_FileID(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.FileID(); got != tt.want {
				t.Errorf("File.FileID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_UniqueID(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.UniqueID(); got != tt.want {
				t.Errorf("File.UniqueID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Atime(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Atime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.Atime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Ctime(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Ctime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.Ctime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Links(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Links(); got != tt.want {
				t.Errorf("File.Links() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_UID(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.UID(); got != tt.want {
				t.Errorf("File.UID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_GID(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.GID(); got != tt.want {
				t.Errorf("File.GID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Type(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Type(); got != tt.want {
				t.Errorf("File.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Reset(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			f.Reset()
		})
	}
}

func TestFile_Open(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if err := f.Open(); (err != nil) != tt.wantErr {
				t.Errorf("File.Open() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_Reopen(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if err := f.Reopen(); (err != nil) != tt.wantErr {
				t.Errorf("File.Reopen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_Read(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	type args struct {
		n uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if err := f.Read(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("File.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_EOF(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.EOF(); got != tt.want {
				t.Errorf("File.EOF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Err(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if err := f.Err(); (err != nil) != tt.wantErr {
				t.Errorf("File.Err() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_Sum(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Sum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Hash(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Hash(); got != tt.want {
				t.Errorf("File.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Hex(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Hex(); got != tt.want {
				t.Errorf("File.Hex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Pos(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Pos(); got != tt.want {
				t.Errorf("File.Pos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Opened(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.Opened(); got != tt.want {
				t.Errorf("File.Opened() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Close(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if err := f.Close(); (err != nil) != tt.wantErr {
				t.Errorf("File.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_ResetHash(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			f.ResetHash()
		})
	}
}

func TestFile_String(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			if got := f.String(); got != tt.want {
				t.Errorf("File.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_MarshalJSON(t *testing.T) {
	type fields struct {
		path     string
		size     uint64
		mode     os.FileMode
		mtime    time.Time
		volumeID uint64
		fileID   uint64
		atime    time.Time
		ctime    time.Time
		nlinks   uint64
		uid      uint64
		gid      uint64
		h        hash.Hash
		sum      []byte
		hash     string
		fh       *os.File
		pos      uint64
		eof      bool
		err      error
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				path:     tt.fields.path,
				size:     tt.fields.size,
				mode:     tt.fields.mode,
				mtime:    tt.fields.mtime,
				volumeID: tt.fields.volumeID,
				fileID:   tt.fields.fileID,
				atime:    tt.fields.atime,
				ctime:    tt.fields.ctime,
				nlinks:   tt.fields.nlinks,
				uid:      tt.fields.uid,
				gid:      tt.fields.gid,
				h:        tt.fields.h,
				sum:      tt.fields.sum,
				hash:     tt.fields.hash,
				fh:       tt.fields.fh,
				pos:      tt.fields.pos,
				eof:      tt.fields.eof,
				err:      tt.fields.err,
			}
			got, err := f.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("File.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
