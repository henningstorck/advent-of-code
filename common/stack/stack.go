package stack

func Push[T any](stack []T, value T) []T {
	return append(stack, value)
}

func Prepend[T any](stack []T, value T) []T {
	return append([]T{value}, stack...)
}

func Pop[T any](stack []T, defaultValue T) ([]T, T, bool) {
	if len(stack) == 0 {
		return stack, defaultValue, false
	} else {
		i := len(stack) - 1
		value := (stack)[i]
		stack = (stack)[:i]
		return stack, value, true
	}
}

func Peek[T any](stack []T, defaultValue T) (T, bool) {
	_, value, ok := Pop(stack, defaultValue)
	return value, ok
}
