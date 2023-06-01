package main

import (
	"chat/pkg"
	"fmt"
)

func main() {
	s1 := []rune("每个 rune 字符和索引在 for-range 循环中是一一对应的")
	s2 := []rune("ge 循环")
	res := pkg.FindSliceInSlice(s1, s2)
	fmt.Println(res, "-------", string(s1[res+1:]))
}
