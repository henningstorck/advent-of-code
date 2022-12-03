package day03

import (
	"strings"

	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/henningstorck/advent-of-code/common/inputreader"
)

func Solve(reader *inputreader.InputReader) (int, int) {
	backpacks := reader.ReadLines(2022, 3)
	return SolvePart1(backpacks), SolvePart2(backpacks)
}

func SolvePart1(backpacks []string) int {
	return functional.Reduce(backpacks, func(sum int, backpack string) int {
		compartment1, compartment2 := SplitIntoCompartments(backpack)
		commonItem := FindCommonItem(compartment1, compartment2)
		priority := GetPriorityFor(commonItem)
		return sum + priority
	}, 0)
}

func SolvePart2(backpacks []string) int {
	sum := 0

	for i := 0; i < len(backpacks); i += 3 {
		commonItem := FindCommonItem(backpacks[i], backpacks[i+1], backpacks[i+2])
		priority := GetPriorityFor(commonItem)
		sum += priority
	}

	return sum
}

func SplitIntoCompartments(backpack string) (string, string) {
	length := len(backpack)
	return backpack[:length/2], backpack[length/2:]
}

func FindCommonItem(reference string, comparisons ...string) rune {
	for _, item := range reference {
		matching := true

		for _, comparison := range comparisons {
			if !strings.ContainsRune(comparison, item) {
				matching = false
			}
		}

		if matching {
			return item
		}
	}

	return rune(0)
}

func GetPriorityFor(item rune) int {
	priority := int(item)

	if priority >= 65 && priority <= 90 {
		return priority - 38
	} else {
		return priority - 96
	}
}
