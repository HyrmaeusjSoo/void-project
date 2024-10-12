package composite

// 转换为指定类型
func Convert[T any](v any) T {
	switch v.(type) {
	case T:
		return v.(T)
	}
	return *new(T)
}
