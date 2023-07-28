package composite

// 对比切片
func CompareSlice[T comparable](s1, s2 []T) bool {
	for k := range s1 {
		if s1[k] != s2[k] {
			return false
		}
	}
	return true
}

// 在长切片中查找短切片
func FindSliceInSlice[T comparable](long, short []T) int {
	si := 0
	for k, v := range long {
		if v == short[si] {
			si++
		} else {
			si = 0
		}
		if si == len(short) {
			return k
		}
	}
	return -1
}
