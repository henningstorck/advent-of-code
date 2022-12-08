package day08_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day08"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	lines := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	resultPart1 := day08.SolvePart1(lines)
	resultPart2 := day08.SolvePart2(lines)
	assert.Equal(t, 21, resultPart1)
	assert.Equal(t, 8, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day08.Solve(reader)
	assert.Equal(t, 1840, resultPart1)
	assert.Equal(t, 405769, resultPart2)
}
