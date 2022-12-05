package stack_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/stack"
	"github.com/stretchr/testify/assert"
)

func TestIsEmptyStack(t *testing.T) {
	stack := stack.Stack{}
	assert.True(t, stack.IsEmpty())
	stack.Push("x")
	assert.False(t, stack.IsEmpty())
}

func TestPushToStack(t *testing.T) {
	stack := stack.Stack{}
	stack.Push("x")
	stack.Push("y")
	assert.EqualValues(t, []string{"x", "y"}, stack)
}

func TestPrependToStack(t *testing.T) {
	stack := stack.Stack{}
	stack.Prepend("x")
	stack.Prepend("y")
	assert.EqualValues(t, []string{"y", "x"}, stack)
}

func TestPopFromStack(t *testing.T) {
	stack := stack.Stack{}
	stack.Push("x")
	stack.Push("y")
	value, ok := stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, "y", value)
	value, ok = stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, "x", value)
	value, ok = stack.Pop()
	assert.False(t, ok)
	assert.Equal(t, "", value)
}
