package day06

import (
	"github.com/henningstorck/advent-of-code/common/inputreader"
)

func Solve(reader *inputreader.InputReader) (int, int) {
	runes := reader.ReadRunes(2022, 6)
	return SolvePart1(runes), SolvePart2(runes)
}

func SolvePart1(runes []rune) int {
	return findStartPos(runes, 4)
}

func SolvePart2(runes []rune) int {
	return findStartPos(runes, 14)
}

func findStartPos(runes []rune, distinctRunes int) int {
	for i := distinctRunes; i < len(runes); i++ {
		set := make(map[rune]bool)

		for j := 0; j < distinctRunes; j++ {
			set[runes[i+j-distinctRunes]] = true
		}

		if len(set) == distinctRunes {
			return i
		}
	}

	return 0
}
