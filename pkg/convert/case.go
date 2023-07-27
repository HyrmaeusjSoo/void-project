package convert

import "strings"

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
