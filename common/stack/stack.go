package stack

type Stack[T any] []T

func (stack Stack[T]) Push(value T) Stack[T] {
	return append(stack, value)
}

func (stack Stack[T]) Prepend(value T) Stack[T] {
	return append([]T{value}, stack...)
}

func (stack Stack[T]) Pop() (Stack[T], T, bool) {
	if len(stack) == 0 {
		var defaultValue T
		return stack, defaultValue, false
	} else {
		i := len(stack) - 1
		value := (stack)[i]
		stack = (stack)[:i]
		return stack, value, true
	}
}

func (stack Stack[T]) Peek() (T, bool) {
	_, value, ok := stack.Pop()
	return value, ok
}
