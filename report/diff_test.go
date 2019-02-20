package report

import (
	"reflect"
	"testing"

	"github.com/rasa/godupless/dupless"
)

func TestNewDiffReport(t *testing.T) {
	type args struct {
		d *dupless.Dupless
	}
	tests := []struct {
		name  string
		args  args
		wantR *DiffReport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := NewDiffReport(tt.args.d); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("NewDiffReport() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestDiffReport_Run(t *testing.T) {
	type fields struct {
		d *dupless.Dupless
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DiffReport{
				d: tt.fields.d,
			}
			r.Run()
		})
	}
}
