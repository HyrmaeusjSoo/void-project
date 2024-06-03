package jwt

import (
	"testing"
)

func BenchmarkGenerateToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateToken(5)
	}
}

func BenchmarkParseToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJpc3MiOiJIeXJtYWV1c2pTb28iLCJzdWIiOiJ2b2lkLXByb2plY3QiLCJhdWQiOlsiNSJdLCJleHAiOjE3MTUxNjA4NjQsIm5iZiI6MTcxNTE1MzY2NCwiaWF0IjoxNzE1MTUzNjY0fQ.-JgDokG_UNm1aJ2hvxA6fOsT30r9FUQycJp7j63BEqU")
	}
}
