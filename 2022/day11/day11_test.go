package day11_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/2022/day11"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	chunks := [][]string{
		{
			"Monkey 0:",
			"  Starting items: 79, 98",
			"  Operation: new = old * 19",
			"  Test: divisible by 23",
			"    If true: throw to monkey 2",
			"    If false: throw to monkey 3",
		},
		{
			"Monkey 1:",
			"  Starting items: 54, 65, 75, 74",
			"  Operation: new = old + 6",
			"  Test: divisible by 19",
			"    If true: throw to monkey 2",
			"    If false: throw to monkey 0",
		},
		{
			"Monkey 2:",
			"  Starting items: 79, 60, 97",
			"  Operation: new = old * old",
			"  Test: divisible by 13",
			"    If true: throw to monkey 1",
			"    If false: throw to monkey 3",
		},
		{
			"Monkey 3:",
			"  Starting items: 74",
			"  Operation: new = old + 3",
			"  Test: divisible by 17",
			"    If true: throw to monkey 0",
			"    If false: throw to monkey 1",
		},
	}

	monkeys := day11.Parse(chunks)
	resultPart1 := day11.SolvePart1(monkeys)
	resultPart2 := day11.SolvePart2(monkeys)
	assert.Equal(t, 10605, resultPart1)
	assert.Equal(t, 2713310158, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := day11.Solve(reader)
	assert.Equal(t, 98280, resultPart1)
	assert.Equal(t, 17673687232, resultPart2)
}
