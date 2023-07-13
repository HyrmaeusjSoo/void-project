package convert

import "strconv"

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
