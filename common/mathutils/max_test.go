package mathutils_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/mathutils"
	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
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
			82,
		},
	}

	for _, test := range tests {
		result := mathutils.Max(test.input...)
		assert.Equal(t, test.expected, result)
	}
}
