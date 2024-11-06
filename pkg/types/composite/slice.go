package composite

// 在切片中查找短切片
func SearchSubSlice[T comparable](haystack, needle []T) int {
	if len(haystack) < len(needle) {
		return -1
	}
	i := 0
	for k, v := range haystack {
		if v == needle[i] {
			i++
		} else {
			i = 0
		}
		if i == len(needle) {
			return k
		}
	}
	return -1
}

// 在切片中向后查找短切片
func SearchSubSliceBackward[T comparable](haystack, needle []T) int {
	if len(haystack) < len(needle) {
		return -1
	}
	e := len(needle)
	match := false
	for i := len(haystack) - 1; i >= 0; i-- {
		if haystack[i] != needle[e-1] {
			if e != len(needle) && haystack[i] == needle[e] {
				e = len(needle) - 1
				match = true
				continue
			}
			e = len(needle)
			match = false
			continue
		}
		e--
		match = true
		if match && e == 0 {
			return i
		}
	}
	return -1
}
