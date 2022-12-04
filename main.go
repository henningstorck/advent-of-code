package main

import (
	"log"

	"github.com/henningstorck/advent-of-code/2022/day04"
	"github.com/henningstorck/advent-of-code/common/inputreader"
)

func main() {
	reader := inputreader.NewInputReader(".")
	resultPart1, resultPart2 := day04.Solve(reader)
	log.Printf("Part 1: %v\n", resultPart1)
	log.Printf("Part 2: %v\n", resultPart2)
}
