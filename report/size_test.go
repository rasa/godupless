package report

import (
	"reflect"
	"testing"

	"github.com/rasa/godupless/dupless"
)

func TestNewSizeReport(t *testing.T) {
	type args struct {
		d *dupless.Dupless
	}
	tests := []struct {
		name  string
		args  args
		wantR *SizeReport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := NewSizeReport(tt.args.d); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("NewSizeReport() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestSizeReport_Run(t *testing.T) {
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
			r := &SizeReport{
				d: tt.fields.d,
			}
			r.Run()
		})
	}
}
