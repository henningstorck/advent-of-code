package queue_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/queue"
	"github.com/stretchr/testify/assert"
)

func TestEnqueueDequeue(t *testing.T) {
	data := queue.Queue[string]{}
	data = data.Enqueue("a")
	data = data.Enqueue("b")
	assert.EqualValues(t, []string{"a", "b"}, data)
	data, value, ok := data.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, "a", value)
	data, value, ok = data.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, "b", value)
	_, value, ok = data.Dequeue()
	assert.False(t, ok)
	assert.Equal(t, "", value)
}
