package day06

import (
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/henningstorck/advent-of-code/common/set"
)

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	runes := reader.ReadRunes(2022, 6, filename)
	return solvePart1(runes), solvePart2(runes)
}

func solvePart1(runes []rune) int {
	return findStartPos(runes, 4)
}

func solvePart2(runes []rune) int {
	return findStartPos(runes, 14)
}

func findStartPos(runes []rune, distinctRunes int) int {
	for i := distinctRunes; i < len(runes); i++ {
		uniqueRunes := set.Set[rune]{}

		for j := 0; j < distinctRunes; j++ {
			uniqueRunes = uniqueRunes.Add(runes[i+j-distinctRunes])
		}

		if len(uniqueRunes) == distinctRunes {
			return i
		}
	}

	return 0
}
