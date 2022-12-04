package day02_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day02"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	lines := []string{
		"A Y",
		"B X",
		"C Z",
	}

	resultPart1 := day02.SolvePart1(lines)
	resultPart2 := day02.SolvePart2(lines)
	assert.Equal(t, 15, resultPart1)
	assert.Equal(t, 12, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day02.Solve(reader)
	assert.Equal(t, 14375, resultPart1)
	assert.Equal(t, 10274, resultPart2)
}
