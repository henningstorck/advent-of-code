package main

import (
	"log"

	"github.com/henningstorck/advent-of-code/2022/day11"
	"github.com/henningstorck/advent-of-code/common/inputreader"
)

func main() {
	reader := inputreader.NewInputReader(".")
	result1, result2 := day11.Solve(reader, "input.txt")
	log.Printf("Part 1: %v\n", result1)
	log.Printf("Part 2: %v\n", result2)
}
