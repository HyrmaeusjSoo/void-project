package logger

import (
	"testing"
)

func TestLevelName(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want string
	}{
		{"t1", AllLevel, ""},
		{"t2", DebugLevel, "debug"},
		{"t3", InfoLevel, "info"},
		{"t4", WarnLevel, "warn"},
		{"t5", ErrorLevel, "error"},
		{"t6", SQLLevel, "sql"},
		{"t7", ServerLevel, "server"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Name(); got != tt.want {
				t.Errorf("Level.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevelValue(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want uint8
	}{
		{"t1", AllLevel, 0},
		{"t2", DebugLevel, 1},
		{"t3", InfoLevel, 2},
		{"t4", WarnLevel, 3},
		{"t5", ErrorLevel, 4},
		{"t6", SQLLevel, 5},
		{"t7", ServerLevel, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Value(); got != tt.want {
				t.Errorf("Level.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLevelName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllLevel.Name()
		ServerLevel.Name()
	}
}

func BenchmarkLevelValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllLevel.Value()
		ServerLevel.Value()
	}
}
