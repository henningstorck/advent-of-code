package day01_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day01"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	input := [][]string{
		{
			"1000",
			"2000",
			"3000",
		},
		{
			"4000",
		},
		{
			"5000",
			"6000",
		},
		{
			"7000",
			"8000",
			"9000",
		},
		{
			"10000",
		},
	}

	inventories := day01.Parse(input)
	resultPart1 := day01.SolvePart1(inventories)
	resultPart2 := day01.SolvePart2(inventories)
	assert.Equal(t, 24000, resultPart1)
	assert.Equal(t, 45000, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day01.Solve(reader)
	assert.Equal(t, 71780, resultPart1)
	assert.Equal(t, 212489, resultPart2)
}
