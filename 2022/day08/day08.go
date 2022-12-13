package day08

import (
	"strconv"

	"github.com/henningstorck/advent-of-code/common/inputreader"
)

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	lines := reader.ReadLines(2022, 8, filename)
	return solvePart1(lines), solvePart2(lines)
}

func solvePart1(lines []string) int {
	treesInSight := 0

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			coveredToNorth, _ := getNorthCoverage(lines, x, y)
			coveredToEast, _ := getEastCoverage(lines, x, y)
			coveredToSouth, _ := getSouthCoverage(lines, x, y)
			coveredToWest, _ := getWestCoverage(lines, x, y)

			if !coveredToNorth || !coveredToEast || !coveredToSouth || !coveredToWest {
				treesInSight++
			}
		}
	}

	return treesInSight
}

func solvePart2(lines []string) int {
	bestScenicScore := 0

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			_, treesToNorth := getNorthCoverage(lines, x, y)
			_, treesToEast := getEastCoverage(lines, x, y)
			_, treesToSouth := getSouthCoverage(lines, x, y)
			_, treesToWest := getWestCoverage(lines, x, y)
			scenicScore := treesToNorth * treesToEast * treesToSouth * treesToWest

			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
		}
	}

	return bestScenicScore
}

func getNorthCoverage(lines []string, x, y int) (bool, int) {
	return getCoverage(lines, x, y, 0, -1)
}

func getEastCoverage(lines []string, x, y int) (bool, int) {
	return getCoverage(lines, x, y, 1, 0)
}

func getSouthCoverage(lines []string, x, y int) (bool, int) {
	return getCoverage(lines, x, y, 0, 1)
}

func getWestCoverage(lines []string, x, y int) (bool, int) {
	return getCoverage(lines, x, y, -1, 0)
}

func getCoverage(lines []string, x, y, xMod, yMod int) (bool, int) {
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
