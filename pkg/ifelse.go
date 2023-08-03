package pkg

// 三元运算
func IfElse[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

// 三元运算 - 执行函数
func IfElseFn[T any](condition bool, trueFunc, falseFunc func() T) T {
	if condition {
		return trueFunc()
	}
	return falseFunc()
}

// 链式操作
type ifelse[T any] struct {
	condition bool
	value     T
}

// 链式操作 If 起始方法
func If[T any](condition bool, value T) *ifelse[T] {
	if condition {
		return &ifelse[T]{true, value}
	}
	return &ifelse[T]{}
}

// 链式操作 ElseIf 中段方法
func (ie *ifelse[T]) ElseIf(condition bool, value T) *ifelse[T] {
	if !ie.condition && condition {
		ie.condition = true
		ie.value = value
	}
	return ie
}

// 链式操作 Else 终结方法
func (ie *ifelse[T]) Else(value T) T {
	return IfElse(ie.condition, ie.value, value)
}

// 链式操作 If 起始方法 - 执行函数
func IfFn[T any](condition bool, function func() T) *ifelse[T] {
	if condition {
		return &ifelse[T]{true, function()}
	}
	return &ifelse[T]{}
}

// 链式操作 ElseIf 中段方法 - 执行函数
func (ie *ifelse[T]) ElseIfFn(condition bool, function func() T) *ifelse[T] {
	if !ie.condition && condition {
		ie.condition = true
		ie.value = function()
	}
	return ie
}

// 链式操作 Else 终结方法 - 执行函数
func (ie *ifelse[T]) ElseFn(function func() T) T {
	return IfElse(ie.condition, ie.value, function())
}
