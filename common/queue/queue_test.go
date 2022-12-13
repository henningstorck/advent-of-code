package queue_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/queue"
	"github.com/stretchr/testify/assert"
)

func TestEnqueueDequeue(t *testing.T) {
	slice := []string{}
	slice = queue.Enqueue(slice, "a")
	slice = queue.Enqueue(slice, "b")
	assert.Equal(t, []string{"a", "b"}, slice)
	slice, value, ok := queue.Dequeue(slice, "")
	assert.True(t, ok)
	assert.Equal(t, "a", value)
	slice, value, ok = queue.Dequeue(slice, "")
	assert.True(t, ok)
	assert.Equal(t, "b", value)
	_, value, ok = queue.Dequeue(slice, "")
	assert.False(t, ok)
	assert.Equal(t, "", value)
}
