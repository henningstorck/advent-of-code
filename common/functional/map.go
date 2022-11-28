package functional

func Map[TValue, TResult any](values []TValue, fn func(TValue) TResult) []TResult {
	result := make([]TResult, len(values))

	for i, value := range values {
		result[i] = fn(value)
	}

	return result
}
