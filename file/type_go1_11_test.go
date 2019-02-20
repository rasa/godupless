// +build go1.11

package file

import (
	"hash"
	"os"
	"testing"
	"time"
)

func TestFile__type(t *testing.T) {
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
			if got := f._type(); got != tt.want {
				t.Errorf("File._type() = %v, want %v", got, tt.want)
			}
		})
	}
}
