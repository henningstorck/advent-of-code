package day14

import (
	"strings"

	"github.com/henningstorck/advent-of-code/common/geometry"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/henningstorck/advent-of-code/common/mathutils"
	"github.com/henningstorck/advent-of-code/common/set"
	"github.com/henningstorck/advent-of-code/common/strutils"
)

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	lines := reader.ReadLines(2022, 14, filename)
	return solvePart1(lines), solvePart2(lines)
}

func solvePart1(lines []string) int {
	blocked := set.Set[geometry.Point2D]{}
	blocked, floor := drawRocks(lines, blocked)
	touchedFloor := false
	i := 0

	for !touchedFloor {
		blocked, touchedFloor = spawnAndDrop(blocked, floor)
		i++
	}

	return i - 1
}

func solvePart2(lines []string) int {
	blocked := set.Set[geometry.Point2D]{}
	blocked, floor := drawRocks(lines, blocked)
	i := 0

	for !blocked.Contains(geometry.NewPoint2D(500, 0)) {
		blocked, _ = spawnAndDrop(blocked, floor)
		i++
	}

	return i
}

func drawRocks(lines []string, blocked set.Set[geometry.Point2D]) (set.Set[geometry.Point2D], int) {
	floor := 0

	for _, line := range lines {
		points := strings.Split(line, " -> ")

		for i := 1; i < len(points); i++ {
			point1 := strings.Split(points[i-1], ",")
			point2 := strings.Split(points[i], ",")
			minX := mathutils.Min(strutils.Atoi(point1[0]), strutils.Atoi(point2[0]))
			maxX := mathutils.Max(strutils.Atoi(point1[0]), strutils.Atoi(point2[0]))
			minY := mathutils.Min(strutils.Atoi(point1[1]), strutils.Atoi(point2[1]))
			maxY := mathutils.Max(strutils.Atoi(point1[1]), strutils.Atoi(point2[1]))

			if maxY+2 > floor {
				floor = maxY + 2
			}

			for x := minX; x <= maxX; x++ {
				for y := minY; y <= maxY; y++ {
					blocked = blocked.Add(geometry.NewPoint2D(x, y))
				}
			}
		}
	}

	return blocked, floor
}

func spawnAndDrop(blocked set.Set[geometry.Point2D], floor int) (set.Set[geometry.Point2D], bool) {
	pos := geometry.NewPoint2D(500, 0)

	for {
		newPos := drop(blocked, pos)

		if newPos.Y >= floor {
			blocked = blocked.Add(pos)
			return blocked, true
		}

		if pos == newPos {
			blocked = blocked.Add(pos)
			return blocked, false
		}

		pos = newPos
	}
}

func drop(blocked set.Set[geometry.Point2D], pos geometry.Point2D) geometry.Point2D {
	mods := []geometry.Point2D{
		geometry.NewPoint2D(0, 1),
		geometry.NewPoint2D(-1, 1),
		geometry.NewPoint2D(1, 1),
	}

	for _, mod := range mods {
		newPos := pos.Add(mod)

		if !blocked.Contains(newPos) {
			return newPos
		}
	}

	return pos
}
