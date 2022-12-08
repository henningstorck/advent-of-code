package day08

import (
	"strconv"

	"github.com/henningstorck/advent-of-code/common/inputreader"
)

func Solve(reader *inputreader.InputReader) (int, int) {
	lines := reader.ReadLines(2022, 8)
	return SolvePart1(lines), SolvePart2(lines)
}

func SolvePart1(lines []string) int {
	treesInSight := 0

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			coveredToNorth, _ := isCoveredToNorth(lines, x, y)
			coveredToEast, _ := isCoveredToEast(lines, x, y)
			coveredToSouth, _ := isCoveredToSouth(lines, x, y)
			coveredToWest, _ := isCoveredToWest(lines, x, y)

			if !coveredToNorth || !coveredToEast || !coveredToSouth || !coveredToWest {
				treesInSight++
			}
		}
	}

	return treesInSight
}

func SolvePart2(lines []string) int {
	bestScenicScore := 0

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			_, treesToNorth := isCoveredToNorth(lines, x, y)
			_, treesToEast := isCoveredToEast(lines, x, y)
			_, treesToSouth := isCoveredToSouth(lines, x, y)
			_, treesToWest := isCoveredToWest(lines, x, y)
			scenicScore := treesToNorth * treesToEast * treesToSouth * treesToWest

			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
		}
	}

	return bestScenicScore
}

func isCoveredToNorth(lines []string, x, y int) (bool, int) {
	return isCovered(lines, x, y, 0, -1)
}

func isCoveredToEast(lines []string, x, y int) (bool, int) {
	return isCovered(lines, x, y, 1, 0)
}

func isCoveredToSouth(lines []string, x, y int) (bool, int) {
	return isCovered(lines, x, y, 0, 1)
}

func isCoveredToWest(lines []string, x, y int) (bool, int) {
	return isCovered(lines, x, y, -1, 0)
}

func isCovered(lines []string, x, y, xMod, yMod int) (bool, int) {
	tree := getTreeAt(lines, x, y)
	treesInSight := 0
	x += xMod
	y += yMod

	for x >= 0 && x < len(lines[0]) && y >= 0 && y < len(lines) {
		treesInSight++
		otherTree := getTreeAt(lines, x, y)

		if otherTree >= tree {
			return true, treesInSight
		}

		x += xMod
		y += yMod
	}

	return false, treesInSight
}

func getTreeAt(lines []string, x, y int) int {
	tree, err := strconv.Atoi(string(lines[y][x]))

	if err != nil {
		return 0
	}

	return tree
}
