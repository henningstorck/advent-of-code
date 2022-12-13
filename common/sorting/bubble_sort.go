package sorting

func BubbleSort[T any](slice []T, compare func(a, b T) bool) []T {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if compare(slice[j], slice[j+1]) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}
