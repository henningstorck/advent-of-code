package queue

func Enqueue[T any](queue []T, value T) []T {
	return append(queue, value)
}

func Dequeue[T any](queue []T, defaultValue T) ([]T, T, bool) {
	if len(queue) == 0 {
		return queue, defaultValue, false
	} else {
		value := queue[0]
		queue = queue[1:]
		return queue, value, true
	}
}
