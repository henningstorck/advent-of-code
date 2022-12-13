package day09

import (
	"strconv"

	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/henningstorck/advent-of-code/common/inputreader"
)

type Pos struct {
	X int
	Y int
}

type Instruction struct {
	Direction string
	Distance  int
}

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	lines := reader.ReadLines(2022, 9, filename)
	instructions := parse(lines)
	return solvePart1(instructions), solvePart2(instructions)
}

func parse(lines []string) []Instruction {
	return functional.Map(lines, func(line string) Instruction {
		direction := line[0:1]
		distance, _ := strconv.Atoi(line[2:])
		return Instruction{direction, distance}
	})
}

func solvePart1(instructions []Instruction) int {
	return countVisits(instructions, 2)
}

func solvePart2(instructions []Instruction) int {
	return countVisits(instructions, 10)
}

func countVisits(instructions []Instruction, totalSegments int) int {
	segments := make([]Pos, totalSegments)
	visited := make(map[Pos]bool)
	visited[Pos{0, 0}] = true

	for _, instruction := range instructions {
		for i := 0; i < instruction.Distance; i++ {
			for j := 0; j < len(segments); j++ {
				if j == 0 {
					segments[j] = moveHead(segments[j], instruction.Direction)
				} else {
					segments[j] = moveTail(segments[j-1], segments[j])

					if j == len(segments)-1 {
						visited[segments[j]] = true
					}
				}
			}
		}
	}

	return len(visited)
}

func moveHead(head Pos, direction string) Pos {
	switch direction {
	case "U":
		return Pos{head.X, head.Y + 1}
	case "R":
		return Pos{head.X + 1, head.Y}
	case "D":
		return Pos{head.X, head.Y - 1}
	case "L":
		return Pos{head.X - 1, head.Y}
	default:
		return head
	}
}

func moveTail(head, tail Pos) Pos {
	diffX := head.X - tail.X
	diffY := head.Y - tail.Y
	okX := diffX >= -1 && diffX <= 1
	okY := diffY >= -1 && diffY <= 1

	if !okX || !okY {
		if diffX == 0 && diffY > 0 {
			return Pos{tail.X, tail.Y + 1}
		} else if diffX == 0 && diffY < 0 {
			return Pos{tail.X, tail.Y - 1}
		} else if diffX > 0 && diffY == 0 {
			return Pos{tail.X + 1, tail.Y}
		} else if diffX < 0 && diffY == 0 {
			return Pos{tail.X - 1, tail.Y}
		} else if diffX > 0 && diffY > 0 {
			return Pos{tail.X + 1, tail.Y + 1}
		} else if diffX > 0 && diffY < 0 {
			return Pos{tail.X + 1, tail.Y - 1}
		} else if diffX < 0 && diffY > 0 {
			return Pos{tail.X - 1, tail.Y + 1}
		} else if diffX < 0 && diffY < 0 {
			return Pos{tail.X - 1, tail.Y - 1}
		}
	}

	return tail
}
