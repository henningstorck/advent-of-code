package queue

type Queue[T any] []T

func (queue Queue[T]) Enqueue(value T) Queue[T] {
	return append(queue, value)
}

func (queue Queue[T]) Dequeue() (Queue[T], T, bool) {
	if len(queue) == 0 {
		var defaultValue T
		return queue, defaultValue, false
	} else {
		value := queue[0]
		queue = queue[1:]
		return queue, value, true
	}
}
