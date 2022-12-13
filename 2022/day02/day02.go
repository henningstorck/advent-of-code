package day02

import (
	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/henningstorck/advent-of-code/common/inputreader"
)

const (
	rock1    = "A"
	paper1   = "B"
	scissor1 = "C"

	rock2    = "X"
	paper2   = "Y"
	scissor2 = "Z"

	lose = "X"
	draw = "Y"
	win  = "Z"
)

var lookupTable1 = map[string]int{
	rock1 + " " + rock2:       1 + 3,
	rock1 + " " + paper2:      2 + 6,
	rock1 + " " + scissor2:    3 + 0,
	paper1 + " " + rock2:      1 + 0,
	paper1 + " " + paper2:     2 + 3,
	paper1 + " " + scissor2:   3 + 6,
	scissor1 + " " + rock2:    1 + 6,
	scissor1 + " " + paper2:   2 + 0,
	scissor1 + " " + scissor2: 3 + 3,
}

var lookupTable2 = map[string]int{
	rock1 + " " + lose:    3 + 0,
	rock1 + " " + draw:    1 + 3,
	rock1 + " " + win:     2 + 6,
	paper1 + " " + lose:   1 + 0,
	paper1 + " " + draw:   2 + 3,
	paper1 + " " + win:    3 + 6,
	scissor1 + " " + lose: 2 + 0,
	scissor1 + " " + draw: 3 + 3,
	scissor1 + " " + win:  1 + 6,
}

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	rounds := reader.ReadLines(2022, 2, filename)
	return solvePart1(rounds), solvePart2(rounds)
}

func solvePart1(rounds []string) int {
	return getPoints(rounds, lookupTable1)
}

func solvePart2(rounds []string) int {
	return getPoints(rounds, lookupTable2)
}

func getPoints(rounds []string, lookupTable map[string]int) int {
	return functional.Reduce(rounds, func(points int, round string) int {
		return points + lookupTable[round]
	}, 0)
}
