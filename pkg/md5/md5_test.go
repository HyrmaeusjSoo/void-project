package md5

import "testing"

func TestGenerateLower(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"abcdef"}, "e80b5017098950fc58aad83c8c14978e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateLower(tt.args.code); got != tt.want {
				t.Errorf("GenerateLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateUpper(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"abcdef"}, "E80B5017098950FC58AAD83C8C14978E"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateUpper(tt.args.code); got != tt.want {
				t.Errorf("GenerateUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaltPassword(t *testing.T) {
	type args struct {
		pwd  string
		salt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"abcdef", "qwerty"}, "213bafe39c3e1f58f71de811a88b15c3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SaltPassword(tt.args.pwd, tt.args.salt); got != tt.want {
				t.Errorf("SaltPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	type args struct {
		rpwd string
		salt string
		pwd  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{"abcdef", "qwerty", "213bafe39c3e1f58f71de811a88b15c3"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPassword(tt.args.rpwd, tt.args.salt, tt.args.pwd); got != tt.want {
				t.Errorf("CheckPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGenerateLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateLower("abcdef")
	}
}

func BenchmarkGenerateUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateUpper("abcdef")
	}
}

func BenchmarkSaltPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SaltPassword("abcdef", "qwerty")
	}
}

func BenchmarkCheckPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckPassword("abcdef", "qwerty", "213bafe39c3e1f58f71de811a88b15c3")
	}
}
