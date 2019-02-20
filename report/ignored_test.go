package report

import (
	"reflect"
	"testing"

	"github.com/rasa/godupless/dupless"
)

func TestNewIgnoredReport(t *testing.T) {
	type args struct {
		d *dupless.Dupless
	}
	tests := []struct {
		name  string
		args  args
		wantR *IgnoredReport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := NewIgnoredReport(tt.args.d); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("NewIgnoredReport() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestIgnoredReport_Run(t *testing.T) {
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
			r := &IgnoredReport{
				d: tt.fields.d,
			}
			r.Run()
		})
	}
}
