package set

var void = struct{}{}

type Set[T comparable] map[T]struct{}

func (set Set[T]) Add(value T) Set[T] {
	if set.Contains(value) {
		return set
	}

	set[value] = void
	return set
}

func (set Set[T]) Remove(value T) Set[T] {
	delete(set, value)
	return set
}

func (set Set[T]) Contains(value T) bool {
	_, ok := set[value]
	return ok
}

func (set Set[T]) Values() []T {
	values := make([]T, 0, len(set))

	for value := range set {
		values = append(values, value)
	}

	return values
}
