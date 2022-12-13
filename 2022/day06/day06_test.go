package day06_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day06"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		filename  string
		expected1 int
		expected2 int
	}{
		{"example01.txt", 7, 19},
		{"example02.txt", 5, 23},
		{"example03.txt", 6, 23},
		{"example04.txt", 10, 29},
		{"example05.txt", 11, 26},
		{"input.txt", 1100, 2421},
	}

	for _, test := range tests {
		reader := inputreader.NewInputReader("../..")
		result1, result2 := day06.Solve(reader, test.filename)
		assert.Equal(t, test.expected1, result1)
		assert.Equal(t, test.expected2, result2)
	}
}
