package day11_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day11"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		filename  string
		expected1 int
		expected2 int
	}{
		{"example.txt", 10605, 2713310158},
		{"input.txt", 98280, 17673687232},
	}

	for _, test := range tests {
		reader := inputreader.NewInputReader("../..")
		result1, result2 := day11.Solve(reader, test.filename)
		assert.Equal(t, test.expected1, result1)
		assert.Equal(t, test.expected2, result2)
	}
}
