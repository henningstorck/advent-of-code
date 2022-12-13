package day10_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day10"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		filename  string
		expected1 int
		expected2 string
	}{
		{"example01.txt", 0, "#####"},
		{"example02.txt", 13140, "##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....\n"},
		{"input.txt", 16020, "####..##..####.#..#.####..##..#....###..\n#....#..#....#.#..#....#.#..#.#....#..#.\n###..#......#..#..#...#..#..#.#....#..#.\n#....#.....#...#..#..#...####.#....###..\n#....#..#.#....#..#.#....#..#.#....#.#..\n####..##..####..##..####.#..#.####.#..#.\n"},
	}

	for _, test := range tests {
		reader := inputreader.NewInputReader("../..")
		result1, result2 := day10.Solve(reader, test.filename)
		assert.Equal(t, test.expected1, result1)
		assert.Equal(t, test.expected2, result2)
	}
}
