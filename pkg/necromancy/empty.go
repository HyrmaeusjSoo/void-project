package necromancy

import "reflect"

func IsEmpty(p any) bool {
	return reflect.ValueOf(p).IsZero()
}

func NotEmpty(p any) bool {
	return !IsEmpty(p)
}
