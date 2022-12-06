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

func findStartPos(runes []rune, distictRunes int) int {
	for i := distictRunes; i < len(runes); i++ {
		set := make(map[rune]bool)

		for j := 0; j < distictRunes; j++ {
			set[runes[i+j-distictRunes]] = true
		}

		if len(set) == distictRunes {
			return i
		}
	}

	return 0
}
