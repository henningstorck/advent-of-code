package day07_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day07"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	lines := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	resultPart1 := day07.SolvePart1(lines)
	resultPart2 := day07.SolvePart2(lines)
	assert.Equal(t, 95437, resultPart1)
	assert.Equal(t, 24933642, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day07.Solve(reader)
	assert.Equal(t, 1447046, resultPart1)
	assert.Equal(t, 578710, resultPart2)
}
