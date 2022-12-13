package day12

import (
	"errors"

	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/henningstorck/advent-of-code/common/queue"
)

type Pos struct {
	X int
	Y int
}

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	lines := reader.ReadLines(2022, 12, filename)
	return solvePart1(lines), solvePart2(lines)
}

func solvePart1(lines []string) int {
	start := findPosition(lines, 'S')
	end := findPosition(lines, 'E')

	return simulate(lines, start, func(oldHeight, newHeight rune) bool {
		return oldHeight >= newHeight-1
	}, func(pos Pos) bool {
		return pos == end
	})
}

func solvePart2(lines []string) int {
	start := findPosition(lines, 'E')

	return simulate(lines, start, func(oldHeight, newHeight rune) bool {
		return oldHeight-1 <= newHeight
	}, func(pos Pos) bool {
		return getHeight(lines, pos) == 'a'
	})
}

func simulate(lines []string, start Pos, canReach func(rune, rune) bool, reachedEnd func(Pos) bool) int {
	nextToCheck := []Pos{}
	visited := make(map[Pos]int)
	nextToCheck = queue.Enqueue(nextToCheck, start)
	visited[start] = 1

	for len(nextToCheck) > 0 {
		var pos Pos
		nextToCheck, pos, _ = queue.Dequeue(nextToCheck, Pos{})

		if reachedEnd(pos) {
			return visited[pos] - 1
		}

		neighbours := getNeighbours(lines, pos, canReach)

		for _, neighbour := range neighbours {
			if visited[neighbour] == 0 {
				visited[neighbour] = visited[pos] + 1
				nextToCheck = queue.Enqueue(nextToCheck, neighbour)
			}
		}
	}

	return 0
}

func findPosition(lines []string, marker rune) Pos {
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			value := rune(lines[y][x])

			if marker == value {
				return Pos{x, y}
			}
		}
	}

	return Pos{0, 0}
}

func getNeighbours(lines []string, pos Pos, canReach func(rune, rune) bool) []Pos {
	posMods := []Pos{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	neighbours := []Pos{}

	for _, posMod := range posMods {
		neighbour, err := getNeighbour(lines, pos, posMod, canReach)

		if err == nil {
			neighbours = append(neighbours, neighbour)
		}
	}

	return neighbours
}

func getNeighbour(lines []string, oldPos Pos, posMod Pos, canReach func(rune, rune) bool) (Pos, error) {
	newPos := Pos{
		X: oldPos.X + posMod.X,
		Y: oldPos.Y + posMod.Y,
	}

	if !isWithinBounds(lines, newPos) {
		return Pos{}, errors.New("position is out of bounds")
	}

	oldHeight := getHeight(lines, oldPos)
	newHeight := getHeight(lines, newPos)

	if !canReach(oldHeight, newHeight) {
		return Pos{}, errors.New("position cannot be reached")
	}

	return newPos, nil
}

func isWithinBounds(lines []string, pos Pos) bool {
	return pos.Y >= 0 && pos.Y < len(lines) && pos.X >= 0 && pos.X < len(lines[0])
}

func getHeight(lines []string, pos Pos) rune {
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
