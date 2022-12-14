package mathutils

func Max(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	max := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}

	return max
}
