package day10

import (
	"strings"

	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/henningstorck/advent-of-code/common/strutils"
)

func Solve(reader *inputreader.InputReader, filename string) (int, string) {
	lines := reader.ReadLines(2022, 10, filename)
	return solvePart1(lines), solvePart2(lines)
}

func solvePart1(lines []string) int {
	x := 1
	cycle := 1
	result := 0

	for _, line := range lines {
		if line == "noop" {
			cycle, result = calcTick(x, cycle, result)
		} else if strings.HasPrefix(line, "addx ") {
			cycle, result = calcTick(x, cycle, result)
			cycle, result = calcTick(x, cycle, result)
			x += strutils.Atoi(line[5:])
		}
	}

	return result
}

func solvePart2(lines []string) string {
	x := 1
	cycle := 1
	result := ""

	for _, line := range lines {
		if line == "noop" {
			cycle, result = drawTick(x, cycle, result)
		} else if strings.HasPrefix(line, "addx ") {
			cycle, result = drawTick(x, cycle, result)
			cycle, result = drawTick(x, cycle, result)
			x += strutils.Atoi(line[5:])
		}
	}

	return result
}

func calcTick(x, cycle, result int) (int, int) {
	if (cycle+20)%40 == 0 {
		result += x * cycle
	}

	cycle++
	return cycle, result
}

func drawTick(x, cycle int, result string) (int, string) {
	pos := (cycle - 1) % 40

	if pos == x-1 || pos == x || pos == x+1 {
		result = result + "#"
	} else {
		result = result + "."
	}

	if cycle%40 == 0 {
		result = result + "\n"
	}

	cycle++
	return cycle, result
}
