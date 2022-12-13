package day05_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day05"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		filename  string
		expected1 string
		expected2 string
	}{
		{"example.txt", "CMZ", "MCD"},
		{"input.txt", "TGWSMRBPN", "TZLTLWRNF"},
	}

	for _, test := range tests {
		reader := inputreader.NewInputReader("../..")
		result1, result2 := day05.Solve(reader, test.filename)
		assert.Equal(t, test.expected1, result1)
		assert.Equal(t, test.expected2, result2)
	}
}
