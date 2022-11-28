package functional_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	tests := []struct {
		input        []int
		expected     int
		fn           func(int, int) int
		initialValue int
	}{
		{
			[]int{},
			0,
			func(sum, value int) int { return sum + value },
			0,
		},
		{
			[]int{1, 2, 3, 4, 5},
			15,
			func(sum, value int) int { return sum + value },
			0,
		},
		{
			[]int{1, 2, 3, 4, 5},
			120,
			func(product, value int) int { return product * value },
			1,
		},
	}

	for _, test := range tests {
		result := functional.Reduce(test.input, test.fn, test.initialValue)
		assert.Equal(t, test.expected, result)
	}
}
