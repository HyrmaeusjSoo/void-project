package composite

import "testing"

func TestCompareSlice(t *testing.T) {
	type args struct {
		s1 []int
		s2 []int
	}
	s1 := []int{10, 20, 30, 40, 50}
	s2 := []int{10, 20, 30, 40, 50}
	s3 := []int{10, 20, 41, 40, 50}
	s4 := []int{10, 20, 30, 40, 0}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{s1, s2}, true},
		{"t2", args{s1, s3}, false},
		{"t3", args{s1, s4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareSlice(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CompareSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchSubSlice(t *testing.T) {
	type args struct {
		haystack []int
		needle   []int
	}
	s1 := []int{10, 20, 30, 40, 50}
	s2 := []int{10, 20}
	s3 := []int{40, 50}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{s1, s2}, 1},
		{"t2", args{s1, s3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchSubSlice(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("SearchSubSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCompareSlice(b *testing.B) {
	s1 := []int{10, 20, 30, 40, 50}
	s2 := []int{10, 20, 30, 40, 50}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CompareSlice(s1, s2)
	}
}

func BenchmarkSearchSubSlice(b *testing.B) {
	s1 := []int{10, 20, 30, 40, 50}
	s2 := []int{40, 50}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchSubSlice(s1, s2)
	}
}
