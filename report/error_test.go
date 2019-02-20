package report

import (
	"reflect"
	"testing"

	"github.com/rasa/godupless/dupless"
)

func TestNewErrorReport(t *testing.T) {
	type args struct {
		d *dupless.Dupless
	}
	tests := []struct {
		name  string
		args  args
		wantR *ErrorReport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := NewErrorReport(tt.args.d); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("NewErrorReport() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestErrorReport_Run(t *testing.T) {
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
			r := &ErrorReport{
				d: tt.fields.d,
			}
			r.Run()
		})
	}
}
