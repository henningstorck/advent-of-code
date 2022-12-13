package day04

import (
	"strconv"
	"strings"

	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/henningstorck/advent-of-code/common/inputreader"
)

type Assignment struct {
	From int
	To   int
}

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	lines := reader.ReadLines(2022, 4, filename)
	pairs := parse(lines)
	return solvePart1(pairs), solvePart2(pairs)
}

func parse(lines []string) [][]Assignment {
	return functional.Map(lines, func(line string) []Assignment {
		return functional.Map(strings.Split(line, ","), func(assignmentStr string) Assignment {
			assignmentSlice := strings.Split(assignmentStr, "-")
			from, _ := strconv.Atoi(assignmentSlice[0])
			to, _ := strconv.Atoi(assignmentSlice[1])
			return Assignment{from, to}
		})
	})
}

func solvePart1(pairs [][]Assignment) int {
	return functional.Reduce(pairs, func(sum int, assigmnents []Assignment) int {
		if isFullyContained(assigmnents[0], assigmnents[1]) {
			return sum + 1
		} else {
			return sum
		}
	}, 0)
}

func solvePart2(pairs [][]Assignment) int {
	return functional.Reduce(pairs, func(sum int, assigmnents []Assignment) int {
		if isPartiallyContained(assigmnents[0], assigmnents[1]) {
			return sum + 1
		} else {
			return sum
		}
	}, 0)
}

func isFullyContained(a Assignment, b Assignment) bool {
	return (a.From <= b.From && a.To >= b.To) || (b.From <= a.From && b.To >= a.To)
}

func isPartiallyContained(a Assignment, b Assignment) bool {
	return (a.From <= b.From && b.From <= a.To) || (a.From <= b.From && b.To <= a.To) || (b.From <= a.From && a.From <= b.To) || (b.From <= a.From && a.To <= b.To)
}
