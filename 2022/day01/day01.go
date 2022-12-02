package day01

import (
	"sort"
	"strconv"

	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/henningstorck/advent-of-code/common/inputreader"
)

func Solve(reader *inputreader.InputReader) (int, int) {
	chunks := reader.ReadChunks(2022, 1)
	inventories := Parse(chunks)
	return SolvePart1(inventories), SolvePart2(inventories)
}

func Parse(chunks [][]string) []int {
	return functional.Map(chunks, func(chunk []string) int {
		return functional.Reduce(chunk, func(calories int, snackStr string) int {
			snack, _ := strconv.Atoi(snackStr)
			return calories + snack
		}, 0)
	})
}

func SolvePart1(inventories []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(inventories)))
	return inventories[0]
}

func SolvePart2(inventories []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(inventories)))
	return inventories[0] + inventories[1] + inventories[2]
}
