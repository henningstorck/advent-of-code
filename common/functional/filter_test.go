package functional_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
		fn       func(int) bool
	}{
		{
			[]int{},
			[]int{},
			func(value int) bool { return value < 10 },
		},
		{
			[]int{1, 5, 10, 15, 20},
			[]int{1, 5},
			func(value int) bool { return value < 10 },
		},
	}

	for _, test := range tests {
		result := functional.Filter(test.input, test.fn)
		assert.Equal(t, test.expected, result)
	}
}
