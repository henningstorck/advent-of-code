package day07

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/henningstorck/advent-of-code/common/inputreader"
)

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	lines := reader.ReadLines(2022, 7, filename)
	return solvePart1(lines), solvePart2(lines)
}

func solvePart1(lines []string) int {
	files, dirs := getFileStats(lines)
	dirs = calcDirSizes(files, dirs)
	sum := 0

	for _, dirSize := range dirs {
		if dirSize <= 100000 {
			sum += dirSize
		}
	}

	return sum
}

func solvePart2(lines []string) int {
	files, dirs := getFileStats(lines)
	dirs = calcDirSizes(files, dirs)
	totalSpace := 70000000
	updateSpace := 30000000
	usedSpace := dirs[""]
	unusedSpace := totalSpace - usedSpace
	requiredSpace := updateSpace - unusedSpace
	candidate := totalSpace

	for _, dirSize := range dirs {
		if dirSize-requiredSpace >= 0 && dirSize-requiredSpace < candidate-requiredSpace {
			candidate = dirSize
		}
	}

	return candidate
}

func getFileStats(lines []string) (map[string]int, map[string]int) {
	currentDir := []string{}
	files := make(map[string]int)
	dirs := make(map[string]int)
	dirs[""] = -1
	fileRegex := regexp.MustCompile(`(\d+) ([a-zA-Z0-9.]+)`)

	for _, line := range lines {
		if line == "$ cd /" {
			currentDir = []string{}
		} else if line == "$ cd .." {
			currentDir = currentDir[:len(currentDir)-1]
		} else if strings.HasPrefix(line, "$ cd ") {
			currentDir = append(currentDir, line[5:])
		} else if strings.HasPrefix(line, "dir ") {
			dirname := line[4:]
			fullPath := getFullPath(currentDir, dirname)
			dirs[fullPath] = -1
		} else if fileRegex.Match([]byte(line)) {
			results := fileRegex.FindStringSubmatch(line)
			fileSize, _ := strconv.Atoi(results[1])
			filename := results[2]
			fullPath := getFullPath(currentDir, filename)
			files[fullPath] = fileSize
		}
	}

	return files, dirs
}

func getFullPath(currentDir []string, filename string) string {
	return strings.Join(append(currentDir, filename), "/")
}

func calcDirSizes(files, dirs map[string]int) map[string]int {
	for dir := range dirs {
		dirSize := 0

		for file, fileSize := range files {
			if strings.HasPrefix(file, dir) {
				dirSize += fileSize
			}
		}

		dirs[dir] = dirSize
	}

	return dirs
}
