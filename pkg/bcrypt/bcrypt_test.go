package bcrypt

import "testing"

func BenchmarkGeneratePassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePassword("Nr99uJwW")
	}
}

func BenchmarkComparePassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ComparePassword("$2a$10$Df024rPqe4SX7KbYJ2t6de.UP46X1SZNaL/PkySuQK8J5NR1HFl5S", "Nr99uJwW")
	}
}
