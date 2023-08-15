package primitive

import (
	"slices"
	"strconv"
	"strings"
)

// string转int，不能转换也返回0
func StringToInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

// string转float64，不能转换也返回0
func StringToFloat64(str string) float64 {
	v, _ := strconv.ParseFloat(str, 64)
	return v
}

// 蛇形字符串转Pascal字符串
func SnakeToPascal(snake string) string {
	pascal := strings.Split(snake, "_")
	for k, v := range pascal {
		s := []rune(v)
		if len(s) == 0 {
			continue
		}
		if s[0] >= 65 && s[0] <= 90 {
			continue
		}
		if s[0] >= 97 && s[0] <= 122 {
			s[0] = s[0] - 32
		}
		pascal[k] = string(s)
	}
	return strings.Join(pascal, "")
}

// Pascal字符串转蛇形字符串
func PascalToSnake(pascal string) string {
	snake := make([]rune, 0, len(pascal))
	for k, v := range pascal {
		if v >= 65 && v <= 90 {
			if k != 0 {
				snake = append(snake, '_')
			}
			snake = append(snake, v+32)
		} else {
			snake = append(snake, v)
		}
	}
	return string(snake)
}

// abcdef转换为qwerty
func ConvertAbcToQwerty(abc string) string {
	return keyboardLayout(abc, 1)
}

// qwerty转换为abcdef
func ConvertQwertyToAbc(qwerty string) string {
	return keyboardLayout(qwerty, 2)
}

// qwerty键盘布局
func keyboardLayout(s string, typ int) string {
	var (
		k1, k2 = []rune("QWERTYUIOPASDFGHJKLZXCVBNM"), []rune("qwertyuiopasdfghjklzxcvbnm")
		r      = make([]rune, 0, len(s))
	)
	for _, v := range s {
		t := v
		if v >= 65 && v <= 90 {
			if typ == 1 {
				t = k1[v-65]
			} else {
				t = rune(slices.Index(k1, v) + 65)
			}
		} else if v >= 97 && v <= 122 {
			if typ == 1 {
				t = k2[v-97]
			} else {
				t = rune(slices.Index(k2, v) + 97)
			}
		}
		r = append(r, t)
	}

	return string(r)
}
