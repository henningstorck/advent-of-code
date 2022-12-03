package day03_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day03"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	input := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	resultPart1 := day03.SolvePart1(input)
	resultPart2 := day03.SolvePart2(input)
	assert.Equal(t, 157, resultPart1)
	assert.Equal(t, 70, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day03.Solve(reader)
	assert.Equal(t, 8088, resultPart1)
	assert.Equal(t, 2522, resultPart2)
}

func TestSplitIntoCompartments(t *testing.T) {
	compartment1, compartment2 := day03.SplitIntoCompartments("abcdef")
	assert.Equal(t, "abc", compartment1)
	assert.Equal(t, "def", compartment2)
}

func TestFindCommonItem(t *testing.T) {
	assert.Equal(t, 'c', day03.FindCommonItem("abc", "dec"))
	assert.Equal(t, 'c', day03.FindCommonItem("abc", "abc", "dec"))
	assert.Equal(t, rune(0), day03.FindCommonItem("abc", "def"))
}

func TestGetPriorityFor(t *testing.T) {
	assert.Equal(t, 1, day03.GetPriorityFor('a'))
	assert.Equal(t, 26, day03.GetPriorityFor('z'))
	assert.Equal(t, 27, day03.GetPriorityFor('A'))
	assert.Equal(t, 52, day03.GetPriorityFor('Z'))
}
