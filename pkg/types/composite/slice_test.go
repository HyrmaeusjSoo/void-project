package composite

import "testing"

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

func BenchmarkSearchSubSlice(b *testing.B) {
	s1 := []int{10, 20, 30, 40, 50}
	s2 := []int{40, 50}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchSubSlice(s1, s2)
	}
}
