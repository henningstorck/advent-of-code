package functional

func Reduce[TValue, TResult any](values []TValue, fn func(TResult, TValue) TResult, initialValue TResult) TResult {
	result := initialValue

	for _, value := range values {
		result = fn(result, value)
	}

	return result
}
