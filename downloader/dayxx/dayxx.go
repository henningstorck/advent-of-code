package dayxx

import (
	"github.com/henningstorck/advent-of-code/common/inputreader"
)

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	lines := reader.ReadLines(0, 0, filename)
	return solvePart1(lines), solvePart2(lines)
}

func solvePart1(lines []string) int {
	return 0
}

func solvePart2(lines []string) int {
	return 0
}
