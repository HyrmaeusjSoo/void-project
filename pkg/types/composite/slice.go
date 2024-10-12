package composite

// 在切片中查找短切片
func SearchSubSlice[T comparable](haystack, needle []T) int {
	si := 0
	for k, v := range haystack {
		if v == needle[si] {
			si++
		} else {
			si = 0
		}
		if si == len(needle) {
			return k
		}
	}
	return -1
}
