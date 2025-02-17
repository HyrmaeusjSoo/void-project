package necromancy

import "reflect"

func IsEmpty(p any) bool {
	return p == nil || reflect.ValueOf(p).IsZero()
}

func NotEmpty(p any) bool {
	return !IsEmpty(p)
}
