package strutils_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/strutils"
	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"x", 0},
		{"-4", -4},
		{"22", 22},
	}

	for _, test := range tests {
		result := strutils.Atoi(test.input)
		assert.Equal(t, test.expected, result)
	}
}
