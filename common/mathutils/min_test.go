package mathutils_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/mathutils"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{51, 2, 82, 5, 14, 7, 7, 4, 3, 20, 74},
			2,
		},
	}

	for _, test := range tests {
		result := mathutils.Min(test.input...)
		assert.Equal(t, test.expected, result)
	}
}
