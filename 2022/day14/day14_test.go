package day14_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day14"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		filename  string
		expected1 int
		expected2 int
	}{
		{"example.txt", 24, 93},
		{"input.txt", 737, 28145},
	}

	for _, test := range tests {
		reader := inputreader.NewInputReader("../..")
		result1, result2 := day14.Solve(reader, test.filename)
		assert.Equal(t, test.expected1, result1)
		assert.Equal(t, test.expected2, result2)
	}
}
