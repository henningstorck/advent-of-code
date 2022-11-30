package day01

import (
	"sort"
	"strconv"

	"github.com/henningstorck/advent-of-code/common/inputreader"
)

func Solve(reader *inputreader.InputReader) (int, int) {
	lines := reader.ReadLines(2022, 1)
	inventories := Parse(lines)
	return SolvePart1(inventories), SolvePart2(inventories)
}

func Parse(lines []string) []int {
	inventory := 0
	inventories := []int{}

	for _, line := range lines {
		if len(line) == 0 {
			inventories = append(inventories, inventory)
			inventory = 0
		} else {
			snack, _ := strconv.Atoi(line)
			inventory += snack
		}
	}

	inventories = append(inventories, inventory)
	return inventories
}

func SolvePart1(inventory []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(inventory)))
	return inventory[0]
}

func SolvePart2(inventory []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(inventory)))
	return inventory[0] + inventory[1] + inventory[2]
}
