package day04_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day04"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	lines := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}

	pairs := day04.Parse(lines)
	resultPart1 := day04.SolvePart1(pairs)
	resultPart2 := day04.SolvePart2(pairs)
	assert.Equal(t, 2, resultPart1)
	assert.Equal(t, 4, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day04.Solve(reader)
	assert.Equal(t, 534, resultPart1)
	assert.Equal(t, 841, resultPart2)
}
