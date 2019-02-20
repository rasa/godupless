package types

import (
	"testing"
)

func TestTypes(t *testing.T) {
	t.Logf("TestTypes")
}

func TestUint64Slice_Len(t *testing.T) {
	tests := []struct {
		name string
		p    Uint64Slice
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("Uint64Slice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Slice_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    Uint64Slice
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Uint64Slice.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Slice_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    Uint64Slice
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestUint64Slice_Sort(t *testing.T) {
	tests := []struct {
		name string
		p    Uint64Slice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Sort()
		})
	}
}
