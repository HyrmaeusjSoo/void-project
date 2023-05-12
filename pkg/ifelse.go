package pkg

func IfElse[T any](exp bool, trueValue, falseValue T) T {
	if exp {
		return trueValue
	} else {
		return falseValue
	}
}

func IfElseFunc[T any](exp bool, trueFunc, falseFunc func() T) T {
	if exp {
		return trueFunc()
	} else {
		return falseFunc()
	}
}
