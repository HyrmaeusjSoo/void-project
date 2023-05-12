package main

import (
	"chat/internal/middleware"
	"fmt"
)

func main() {
	// t, err := middleware.GenerateToken(12345)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(t)

	claims, err := middleware.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEyMzQ1LCJpc3MiOiIxMjM0NSIsInN1YiI6ImNoYXQiLCJleHAiOjE2ODY0NTM2MzgsIm5iZiI6MTY4Mzg2MTYzOCwiaWF0IjoxNjgzODYxNjM4fQ.mlPJkflwxqfOVK04W2SqrupnKqv3QesRbJWGB7KT3g0")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(claims.UserID, claims.Issuer)
}
