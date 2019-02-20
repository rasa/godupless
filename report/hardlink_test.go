package report

import (
	"reflect"
	"testing"

	"github.com/rasa/godupless/dupless"
)

func TestNewHardLinkReport(t *testing.T) {
	type args struct {
		d *dupless.Dupless
	}
	tests := []struct {
		name  string
		args  args
		wantR *HardLinkReport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := NewHardLinkReport(tt.args.d); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("NewHardLinkReport() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestHardLinkReport_Run(t *testing.T) {
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
			r := &HardLinkReport{
				d: tt.fields.d,
			}
			r.Run()
		})
	}
}
