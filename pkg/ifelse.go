package pkg

func IfElse[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
func IfElseFn[T any](condition bool, trueFunc, falseFunc func() T) T {
	if condition {
		return trueFunc()
	}
	return falseFunc()
}

type ifelse[T any] struct {
	condition bool
	value     T
}

func If[T any](condition bool, value T) *ifelse[T] {
	if condition {
		return &ifelse[T]{true, value}
	}
	return &ifelse[T]{}
}
func (ie *ifelse[T]) ElseIf(condition bool, value T) *ifelse[T] {
	if !ie.condition && condition {
		ie.condition = true
		ie.value = value
	}
	return ie
}
func (ie *ifelse[T]) Else(value T) T {
	return IfElse(ie.condition, ie.value, value)
}

func IfFn[T any](condition bool, function func() T) *ifelse[T] {
	if condition {
		return &ifelse[T]{true, function()}
	}
	return &ifelse[T]{}
}
func (ie *ifelse[T]) ElseIfFn(condition bool, function func() T) *ifelse[T] {
	if !ie.condition && condition {
		ie.condition = true
		ie.value = function()
	}
	return ie
}
func (ie *ifelse[T]) ElseFn(function func() T) T {
	return IfElse(ie.condition, ie.value, function())
}
