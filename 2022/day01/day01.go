package day01

import (
	"sort"
	"strconv"

	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/henningstorck/advent-of-code/common/inputreader"
)

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	chunks := reader.ReadChunks(2022, 1, filename)
	inventories := parse(chunks)
	return solvePart1(inventories), solvePart2(inventories)
}

func parse(chunks [][]string) []int {
	return functional.Map(chunks, func(chunk []string) int {
		return functional.Reduce(chunk, func(calories int, snackStr string) int {
			snack, _ := strconv.Atoi(snackStr)
			return calories + snack
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
