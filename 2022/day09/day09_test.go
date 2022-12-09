package day09_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day09"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	tests := []struct {
		lines         []string
		expectedPart1 int
		expectedPart2 int
	}{
		{
			[]string{
				"R 4",
				"U 4",
				"L 3",
				"D 1",
				"R 4",
				"D 1",
				"L 5",
				"R 2",
			},
			13,
			1,
		},
		{
			[]string{
				"R 5",
				"U 8",
				"L 8",
				"D 3",
				"R 17",
				"D 10",
				"L 25",
				"U 20",
			},
			88,
			36,
		},
	}

	for _, test := range tests {
		instructions := day09.Parse(test.lines)
		resultPart1 := day09.SolvePart1(instructions)
		resultPart2 := day09.SolvePart2(instructions)
		assert.Equal(t, test.expectedPart1, resultPart1)
		assert.Equal(t, test.expectedPart2, resultPart2)
	}
}

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day09.Solve(reader)
	assert.Equal(t, 6745, resultPart1)
	assert.Equal(t, 2793, resultPart2)
}
