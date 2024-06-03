package primitive

import (
	"testing"
)

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{"t1", "1", 1},
		{"t2", "a", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt(tt.str); got != tt.want {
				t.Errorf("StringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToFloat64(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want float64
	}{
		{"t1", "1", 1.00},
		{"t2", "2.00", 2.00},
		{"t3", "a", 0.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToFloat64(tt.str); got != tt.want {
				t.Errorf("StringToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeToPascal(t *testing.T) {
	tests := []struct {
		name  string
		snake string
		want  string
	}{
		{"t1", "snake_to_pascel", "SnakeToPascel"},
		{"t2", "Snake_to_Pascel", "SnakeToPascel"},
		{"t3", "sNake_to_paScel", "SNakeToPaScel"},
		{"t4", "snaketopascel", "Snaketopascel"},
		{"t5", "_snake_to_pascel_", "SnakeToPascel"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeToPascal(tt.snake); got != tt.want {
				t.Errorf("SnakeToPascal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPascalToSnake(t *testing.T) {
	tests := []struct {
		name   string
		pascal string
		want   string
	}{
		{"t1", "SnakeToPascel", "snake_to_pascel"},
		{"t3", "SnakeTopascel", "snake_topascel"},
		{"t4", "SnaketopasceL", "snaketopasce_l"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PascalToSnake(tt.pascal); got != tt.want {
				t.Errorf("PascalToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertAbcToQwerty(t *testing.T) {
	tests := []struct {
		name string
		abc  string
		want string
	}{
		{"t1", "abcdef", "qwerty"},
		{"t2", "xyz", "bnm"},
		{"t3", "hello world", "itssg vgksr"},
		{"t4", "AbCdEf", "QwErTy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAbcToQwerty(tt.abc); got != tt.want {
				t.Errorf("ConvertAbcToQwerty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertQwertyToAbc(t *testing.T) {
	tests := []struct {
		name   string
		qwerty string
		want   string
	}{
		{"t1", "qwerty", "abcdef"},
		{"t2", "bnm", "xyz"},
		{"t3", "itssg vgksr", "hello world"},
		{"t4", "QwErTy", "AbCdEf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertQwertyToAbc(tt.qwerty); got != tt.want {
				t.Errorf("ConvertQwertyToAbc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToInt("12345")
	}
}

func BenchmarkStringToFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToFloat64("123.45")
	}
}

func BenchmarkSnakeToPascal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SnakeToPascal("snake_to_pascal")
	}
}

func BenchmarkPascalToSnake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PascalToSnake("PascalToSnake")
	}
}

func BenchmarkConvertAbcToQwerty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertAbcToQwerty("abcdef")
	}
}

func BenchmarkConvertQwertyToAbc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertQwertyToAbc("qwerty")
	}
}
