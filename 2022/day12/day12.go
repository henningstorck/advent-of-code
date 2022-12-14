package day12

import (
	"errors"

	"github.com/henningstorck/advent-of-code/common/geometry"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/henningstorck/advent-of-code/common/queue"
)

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	lines := reader.ReadLines(2022, 12, filename)
	return solvePart1(lines), solvePart2(lines)
}

func solvePart1(lines []string) int {
	start := findPosition(lines, 'S')
	end := findPosition(lines, 'E')

	return simulate(lines, start, func(oldHeight, newHeight rune) bool {
		return oldHeight >= newHeight-1
	}, func(pos geometry.Point2D) bool {
		return pos == end
	})
}

func solvePart2(lines []string) int {
	start := findPosition(lines, 'E')

	return simulate(lines, start, func(oldHeight, newHeight rune) bool {
		return oldHeight-1 <= newHeight
	}, func(pos geometry.Point2D) bool {
		return getHeight(lines, pos) == 'a'
	})
}

func simulate(lines []string, start geometry.Point2D, canReach func(rune, rune) bool, reachedEnd func(geometry.Point2D) bool) int {
	nextToCheck := queue.Queue[geometry.Point2D]{}
	visited := make(map[geometry.Point2D]int)
	nextToCheck = nextToCheck.Enqueue(start)
	visited[start] = 1

	for len(nextToCheck) > 0 {
		var pos geometry.Point2D
		nextToCheck, pos, _ = nextToCheck.Dequeue()

		if reachedEnd(pos) {
			return visited[pos] - 1
		}

		neighbours := getNeighbours(lines, pos, canReach)

		for _, neighbour := range neighbours {
			if visited[neighbour] == 0 {
				visited[neighbour] = visited[pos] + 1
				nextToCheck = nextToCheck.Enqueue(neighbour)
			}
		}
	}

	return 0
}

func findPosition(lines []string, marker rune) geometry.Point2D {
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			value := rune(lines[y][x])

			if marker == value {
				return geometry.NewPoint2D(x, y)
			}
		}
	}

	return geometry.NewPoint2D(0, 0)
}

func getNeighbours(lines []string, pos geometry.Point2D, canReach func(rune, rune) bool) []geometry.Point2D {
	posMods := []geometry.Point2D{
		geometry.NewPoint2D(1, 0),
		geometry.NewPoint2D(0, 1),
		geometry.NewPoint2D(-1, 0),
		geometry.NewPoint2D(0, -1),
	}

	neighbours := []geometry.Point2D{}

	for _, posMod := range posMods {
		neighbour, err := getNeighbour(lines, pos, posMod, canReach)

		if err == nil {
			neighbours = append(neighbours, neighbour)
		}
	}

	return neighbours
}

func getNeighbour(lines []string, oldPos geometry.Point2D, posMod geometry.Point2D, canReach func(rune, rune) bool) (geometry.Point2D, error) {
	newPos := oldPos.Add(posMod)

	if !isWithinBounds(lines, newPos) {
		return geometry.Point2D{}, errors.New("position is out of bounds")
	}

	oldHeight := getHeight(lines, oldPos)
	newHeight := getHeight(lines, newPos)

	if !canReach(oldHeight, newHeight) {
		return geometry.Point2D{}, errors.New("position cannot be reached")
	}

	return newPos, nil
}

func isWithinBounds(lines []string, pos geometry.Point2D) bool {
	return pos.Y >= 0 && pos.Y < len(lines) && pos.X >= 0 && pos.X < len(lines[0])
}

func getHeight(lines []string, pos geometry.Point2D) rune {
	height := rune(lines[pos.Y][pos.X])

	switch height {
	case 'S':
		return 'a'
	case 'E':
		return 'z'
	default:
		return height
	}
}
