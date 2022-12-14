package day09

import (
	"strconv"

	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/henningstorck/advent-of-code/common/geometry"
	"github.com/henningstorck/advent-of-code/common/inputreader"
)

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
	segments := make([]geometry.Point2D, totalSegments)
	visited := make(map[geometry.Point2D]bool)
	visited[geometry.NewPoint2D(0, 0)] = true

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

func moveHead(head geometry.Point2D, direction string) geometry.Point2D {
	switch direction {
	case "U":
		return geometry.NewPoint2D(head.X, head.Y+1)
	case "R":
		return geometry.NewPoint2D(head.X+1, head.Y)
	case "D":
		return geometry.NewPoint2D(head.X, head.Y-1)
	case "L":
		return geometry.NewPoint2D(head.X-1, head.Y)
	default:
		return head
	}
}

func moveTail(head, tail geometry.Point2D) geometry.Point2D {
	diffX := head.X - tail.X
	diffY := head.Y - tail.Y
	okX := diffX >= -1 && diffX <= 1
	okY := diffY >= -1 && diffY <= 1

	if !okX || !okY {
		if diffX == 0 && diffY > 0 {
			return tail.AddXY(0, 1)
		} else if diffX == 0 && diffY < 0 {
			return tail.AddXY(0, -1)
		} else if diffX > 0 && diffY == 0 {
			return tail.AddXY(1, 0)
		} else if diffX < 0 && diffY == 0 {
			return tail.AddXY(-1, 0)
		} else if diffX > 0 && diffY > 0 {
			return tail.AddXY(1, 1)
		} else if diffX > 0 && diffY < 0 {
			return tail.AddXY(1, -1)
		} else if diffX < 0 && diffY > 0 {
			return tail.AddXY(-1, 1)
		} else if diffX < 0 && diffY < 0 {
			return tail.AddXY(-1, -1)
		}
	}

	return tail
}
