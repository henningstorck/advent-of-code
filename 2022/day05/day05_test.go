package day05_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day05"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	chunks := [][]string{
		{
			"    [D]    ",
			"[N] [C]    ",
			"[Z] [M] [P]",
			" 1   2   3 ",
		},
		{
			"move 1 from 2 to 1",
			"move 3 from 1 to 3",
			"move 2 from 2 to 1",
			"move 1 from 1 to 2",
		},
	}

	piles, instructions := day05.Parse(chunks)
	resultPart1 := day05.SolvePart1(piles, instructions)
	resultPart2 := day05.SolvePart2(piles, instructions)
	assert.Equal(t, "CMZ", resultPart1)
	assert.Equal(t, "MCD", resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day05.Solve(reader)
	assert.Equal(t, "TGWSMRBPN", resultPart1)
	assert.Equal(t, "TZLTLWRNF", resultPart2)
}
