// +build windows

package file

import (
	"hash"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_loadVolumeMap(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := loadVolumeMap(); (err != nil) != tt.wantErr {
				t.Errorf("loadVolumeMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_VolumeName(t *testing.T) {
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
		name       string
		fields     fields
		wantVolume string
		wantErr    bool
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
			gotVolume, err := f.VolumeName()
			if (err != nil) != tt.wantErr {
				t.Errorf("File.VolumeName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotVolume != tt.wantVolume {
				t.Errorf("File.VolumeName() = %v, want %v", gotVolume, tt.wantVolume)
			}
		})
	}
}

func TestGetVolumes(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetVolumes()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVolumes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVolumes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsVolume(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsVolume(tt.args.path); got != tt.want {
				t.Errorf("IsVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}
