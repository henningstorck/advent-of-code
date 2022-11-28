package functional

func Filter[TValue any](values []TValue, fn func(TValue) bool) []TValue {
	result := make([]TValue, 0)

	for _, value := range values {
		if fn(value) {
			result = append(result, value)
		}
	}

	return result
}
