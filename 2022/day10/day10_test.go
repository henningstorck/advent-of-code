package day10_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day10"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day10.Solve(reader)
	assert.Equal(t, 16020, resultPart1)
	assert.Equal(t, "####..##..####.#..#.####..##..#....###..\n#....#..#....#.#..#....#.#..#.#....#..#.\n###..#......#..#..#...#..#..#.#....#..#.\n#....#.....#...#..#..#...####.#....###..\n#....#..#.#....#..#.#....#..#.#....#.#..\n####..##..####..##..####.#..#.####.#..#.\n", resultPart2)
}
