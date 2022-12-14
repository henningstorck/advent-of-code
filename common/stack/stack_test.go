package stack_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/stack"
	"github.com/stretchr/testify/assert"
)

func TestPushToStack(t *testing.T) {
	data := stack.Stack[string]{}
	data = data.Push("x")
	data = data.Push("y")
	assert.EqualValues(t, []string{"x", "y"}, data)
}

func TestPrependToStack(t *testing.T) {
	data := stack.Stack[string]{}
	data = data.Prepend("x")
	data = data.Prepend("y")
	assert.EqualValues(t, []string{"y", "x"}, data)
}

func TestPopFromStack(t *testing.T) {
	data := stack.Stack[string]{}
	data = data.Push("x")
	data = data.Push("y")
	data, value, ok := data.Pop()
	assert.True(t, ok)
	assert.Equal(t, "y", value)
	data, value, ok = data.Pop()
	assert.True(t, ok)
	assert.Equal(t, "x", value)
	_, value, ok = data.Pop()
	assert.False(t, ok)
	assert.Equal(t, "", value)
}
