package day06_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day06"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	tests := []struct {
		runes         []rune
		expectedPart1 int
		expectedPart2 int
	}{
		{
			runes:         []rune{'m', 'j', 'q', 'j', 'p', 'q', 'm', 'g', 'b', 'l', 'j', 's', 'p', 'h', 'd', 'z', 't', 'n', 'v', 'j', 'f', 'q', 'w', 'r', 'c', 'g', 's', 'm', 'l', 'b'},
			expectedPart1: 7,
			expectedPart2: 19,
		},
		{
			runes:         []rune{'b', 'v', 'w', 'b', 'j', 'p', 'l', 'b', 'g', 'v', 'b', 'h', 's', 'r', 'l', 'p', 'g', 'd', 'm', 'j', 'q', 'w', 'f', 't', 'v', 'n', 'c', 'z'},
			expectedPart1: 5,
			expectedPart2: 23,
		},
		{
			runes:         []rune{'n', 'p', 'p', 'd', 'v', 'j', 't', 'h', 'q', 'l', 'd', 'p', 'w', 'n', 'c', 'q', 's', 'z', 'v', 'f', 't', 'b', 'r', 'm', 'j', 'l', 'h', 'g'},
			expectedPart1: 6,
			expectedPart2: 23,
		},
		{
			runes:         []rune{'n', 'z', 'n', 'r', 'n', 'f', 'r', 'f', 'n', 't', 'j', 'f', 'm', 'v', 'f', 'w', 'm', 'z', 'd', 'f', 'j', 'l', 'v', 't', 'q', 'n', 'b', 'h', 'c', 'p', 'r', 's', 'g'},
			expectedPart1: 10,
			expectedPart2: 29,
		},
		{
			runes:         []rune{'z', 'c', 'f', 'z', 'f', 'w', 'z', 'z', 'q', 'f', 'r', 'l', 'j', 'w', 'z', 'l', 'r', 'f', 'n', 'p', 'q', 'd', 'b', 'h', 't', 'm', 's', 'c', 'g', 'v', 'j', 'w'},
			expectedPart1: 11,
			expectedPart2: 26,
		},
	}

	for _, test := range tests {
		resultPart1 := day06.SolvePart1(test.runes)
		resultPart2 := day06.SolvePart2(test.runes)
		assert.Equal(t, test.expectedPart1, resultPart1)
		assert.Equal(t, test.expectedPart2, resultPart2)
	}
}

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day06.Solve(reader)
	assert.Equal(t, 1100, resultPart1)
	assert.Equal(t, 2421, resultPart2)
}
