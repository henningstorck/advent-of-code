package day01

import (
	"sort"

	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/henningstorck/advent-of-code/common/strutils"
)

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	chunks := reader.ReadChunks(2022, 1, filename)
	inventories := parse(chunks)
	return solvePart1(inventories), solvePart2(inventories)
}

func parse(chunks [][]string) []int {
	return functional.Map(chunks, func(chunk []string) int {
		return functional.Reduce(chunk, func(calories int, snack string) int {
			return calories + strutils.Atoi(snack)
		}, 0)
	})
}

func solvePart1(inventories []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(inventories)))
	return inventories[0]
}

func solvePart2(inventories []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(inventories)))
	return inventories[0] + inventories[1] + inventories[2]
}
