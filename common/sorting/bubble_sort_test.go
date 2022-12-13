package sorting_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/sorting"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{51, 2, 82, 5, 14, 7, 7, 4, 3, 20, 74},
			[]int{2, 3, 4, 5, 7, 7, 14, 20, 51, 74, 82},
		},
	}

	for _, test := range tests {
		sorted := sorting.BubbleSort(test.input, func(a, b int) bool {
			return a > b
		})

		assert.Equal(t, test.expected, sorted)
	}
}
