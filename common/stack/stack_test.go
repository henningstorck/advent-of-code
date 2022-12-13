package stack_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/stack"
	"github.com/stretchr/testify/assert"
)

func TestPushToStack(t *testing.T) {
	slice := []string{}
	slice = stack.Push(slice, "x")
	slice = stack.Push(slice, "y")
	assert.EqualValues(t, []string{"x", "y"}, slice)
}

func TestPrependToStack(t *testing.T) {
	slice := []string{}
	slice = stack.Prepend(slice, "x")
	slice = stack.Prepend(slice, "y")
	assert.EqualValues(t, []string{"y", "x"}, slice)
}

func TestPopFromStack(t *testing.T) {
	slice := []string{}
	slice = stack.Push(slice, "x")
	slice = stack.Push(slice, "y")
	slice, value, ok := stack.Pop(slice, "")
	assert.True(t, ok)
	assert.Equal(t, "y", value)
	slice, value, ok = stack.Pop(slice, "")
	assert.True(t, ok)
	assert.Equal(t, "x", value)
	_, value, ok = stack.Pop(slice, "")
	assert.False(t, ok)
	assert.Equal(t, "", value)
}
