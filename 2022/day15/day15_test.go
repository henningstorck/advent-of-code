package day15_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day15"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		filename  string
		row       int
		expected1 int
		expected2 int
	}{
		{"example.txt", 10, 26, 56000011},
		{"input.txt", 2000000, 5142231, 10884459367718},
	}

	for _, test := range tests {
		reader := inputreader.NewInputReader("../..")
		result1, result2 := day15.Solve(reader, test.filename, test.row)
		assert.Equal(t, test.expected1, result1)
		assert.Equal(t, test.expected2, result2)
	}
}
