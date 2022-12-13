package day12_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day12"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	lines := []string{
		"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}

	resultPart1 := day12.SolvePart1(lines)
	resultPart2 := day12.SolvePart2(lines)
	assert.Equal(t, 31, resultPart1)
	assert.Equal(t, 29, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day12.Solve(reader)
	assert.Equal(t, 447, resultPart1)
	assert.Equal(t, 446, resultPart2)
}
