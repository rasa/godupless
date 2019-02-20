package report

import (
	"reflect"
	"testing"

	"github.com/rasa/godupless/dupless"
)

func TestNewDirReport(t *testing.T) {
	type args struct {
		d *dupless.Dupless
	}
	tests := []struct {
		name  string
		args  args
		wantR *DirReport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := NewDirReport(tt.args.d); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("NewDirReport() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestDirReport_Run(t *testing.T) {
	type fields struct {
		d    *dupless.Dupless
		Dirs map[string]*Dir
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DirReport{
				d:    tt.fields.d,
				Dirs: tt.fields.Dirs,
			}
			r.Run()
		})
	}
}

func TestDirReport_addPath(t *testing.T) {
	type fields struct {
		d    *dupless.Dupless
		Dirs map[string]*Dir
	}
	type args struct {
		path string
		size uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DirReport{
				d:    tt.fields.d,
				Dirs: tt.fields.Dirs,
			}
			r.addPath(tt.args.path, tt.args.size)
		})
	}
}
