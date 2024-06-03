package pkg

import (
	"testing"
)

func TestGetRootPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRootPath(); got != tt.want {
				t.Errorf("GetRootPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubPath(t *testing.T) {
	type args struct {
		dir string
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
			if got := SubPath(tt.args.dir); got != tt.want {
				t.Errorf("SubPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetRootPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRootPath()
	}
}

func BenchmarkSubPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubPath("void-project/pkg")
	}
}
