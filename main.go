package main

import (
	"log"

	"github.com/henningstorck/advent-of-code/2022/day15"
	"github.com/henningstorck/advent-of-code/common/inputreader"
)

func main() {
	reader := inputreader.NewInputReader(".")
	result1, result2 := day15.Solve(reader, "input.txt", 2000000)
	log.Printf("Part 1: %v\n", result1)
	log.Printf("Part 2: %v\n", result2)
}
