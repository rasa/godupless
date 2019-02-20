// +build nacl js,wasm

package file

import (
	"hash"
	"os"
	"testing"
	"time"
)

func TestFile_stat(t *testing.T) {
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
		fi os.FileInfo
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
			if err := f.stat(tt.args.fi); (err != nil) != tt.wantErr {
				t.Errorf("File.stat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
