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

func SolvePart1(inventory []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(inventory)))
	return inventory[0]
}

func SolvePart2(inventory []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(inventory)))
	return inventory[0] + inventory[1] + inventory[2]
}
