package util

import (
	"reflect"
	"syscall"
	"testing"
	"time"
)

func TestUtil(t *testing.T) {
	t.Logf("TestUtil")
}

func TestBasename(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Basename(tt.args.path); got != tt.want {
				t.Errorf("Basename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirname(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dirname(tt.args.path); got != tt.want {
				t.Errorf("Dirname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDump(t *testing.T) {
	type args struct {
		s string
		x interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Dump(tt.args.s, tt.args.x)
		})
	}
}

func TestNormalizePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizePath(tt.args.path); got != tt.want {
				t.Errorf("NormalizePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPause(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pause()
		})
	}
}

func TestTimespecToTime(t *testing.T) {
	type args struct {
		ts syscall.Timespec
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimespecToTime(tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimespecToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimSuffix(t *testing.T) {
	type args struct {
		s      string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimSuffix(tt.args.s, tt.args.suffix); got != tt.want {
				t.Errorf("TrimSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
