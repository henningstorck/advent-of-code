package stack

type Stack []string

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Push(value string) {
	*stack = append(*stack, value)
}

func (stack *Stack) Prepend(value string) {
	*stack = append([]string{value}, *stack...)
}

func (stack *Stack) Pop() (string, bool) {
	if stack.IsEmpty() {
		return "", false
	} else {
		i := len(*stack) - 1
		value := (*stack)[i]
		*stack = (*stack)[:i]
		return value, true
	}
}

func (stack *Stack) Peek() (string, bool) {
	if stack.IsEmpty() {
		return "", false
	} else {
		i := len(*stack) - 1
		value := (*stack)[i]
		return value, true
	}
}
