package day13_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day13"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	chunks := [][]string{
		{
			"[1,1,3,1,1]",
			"[1,1,5,1,1]",
		},
		{
			"[[1],[2,3,4]]",
			"[[1],4]",
		},
		{
			"[9]",
			"[[8,7,6]]",
		},
		{
			"[[4,4],4,4]",
			"[[4,4],4,4,4]",
		},
		{
			"[7,7,7,7]",
			"[7,7,7]",
		},
		{
			"[]",
			"[3]",
		},
		{
			"[[[]]]",
			"[[]]",
		},
		{
			"[1,[2,[3,[4,[5,6,7]]]],8,9]",
			"[1,[2,[3,[4,[5,6,0]]]],8,9]",
		},
	}

	resultPart1 := day13.SolvePart1(chunks)
	resultPart2 := day13.SolvePart2(chunks)
	assert.Equal(t, 13, resultPart1)
	assert.Equal(t, 140, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day13.Solve(reader)
	assert.Equal(t, 5003, resultPart1)
	assert.Equal(t, 20280, resultPart2)
}
