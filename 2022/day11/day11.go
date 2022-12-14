package day11

import (
	"sort"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/henningstorck/advent-of-code/common/strutils"
	"github.com/jinzhu/copier"
)

type Monkey struct {
	Items       []int
	Operation   string
	DivisibleBy int
	IfTrue      int
	IfFalse     int
	Inspections int
}

func Solve(reader *inputreader.InputReader, filename string) (int, int) {
	chunks := reader.ReadChunks(2022, 11, filename)
	monkeys := parse(chunks)
	return solvePart1(monkeys), solvePart2(monkeys)
}

func parse(chunks [][]string) []*Monkey {
	return functional.Map(chunks, func(lines []string) *Monkey {
		monkey := &Monkey{}

		for _, line := range lines {
			if strings.HasPrefix(line, "Monkey ") {
				continue
			} else if strings.HasPrefix(line, "  Starting items: ") {
				startingItemsStr := line[18:]
				startingItemsStrSlice := strings.Split(startingItemsStr, ", ")
				startingItems := functional.Map(startingItemsStrSlice, func(value string) int {
					return strutils.Atoi(value)
				})
				monkey.Items = startingItems
			} else if strings.HasPrefix(line, "  Operation: new = ") {
				operation := line[19:]
				monkey.Operation = operation
			} else if strings.HasPrefix(line, "  Test: divisible by ") {
				divisibleBy := line[21:]
				monkey.DivisibleBy = strutils.Atoi(divisibleBy)
			} else if strings.HasPrefix(line, "    If true: throw to monkey ") {
				ifTrue := line[29:]
				monkey.IfTrue = strutils.Atoi(ifTrue)
			} else if strings.HasPrefix(line, "    If false: throw to monkey ") {
				ifFalse := line[30:]
				monkey.IfFalse = strutils.Atoi(ifFalse)
			}
		}

		return monkey
	})
}

func solvePart1(monkeys []*Monkey) int {
	var monkeysCopy []*Monkey
	copier.CopyWithOption(&monkeysCopy, &monkeys, copier.Option{DeepCopy: true})

	return simulate(monkeysCopy, 20, func(item int) int {
		return item / 3
	})
}

func solvePart2(monkeys []*Monkey) int {
	var monkeysCopy []*Monkey
	copier.CopyWithOption(&monkeysCopy, &monkeys, copier.Option{DeepCopy: true})
	superModulo := getSuperModulo(monkeysCopy)

	return simulate(monkeysCopy, 10000, func(item int) int {
		return item % superModulo
	})
}

func simulate(monkeys []*Monkey, rounds int, decreaseWorryLevel func(int) int) int {
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				monkey.Inspections++
				item = evalOperation(monkey.Operation, item)
				item = decreaseWorryLevel(item)

				if item%monkey.DivisibleBy == 0 {
					monkeys[monkey.IfTrue].Items = append(monkeys[monkey.IfTrue].Items, item)
				} else {
					monkeys[monkey.IfFalse].Items = append(monkeys[monkey.IfFalse].Items, item)
				}
			}

			monkey.Items = []int{}
		}
	}

	return getLevelOfMonkeyBusiness(monkeys)
}

func getSuperModulo(monkeys []*Monkey) int {
	return functional.Reduce(monkeys, func(superModulo int, monkey *Monkey) int {
		return superModulo * monkey.DivisibleBy
	}, 1)
}

func evalOperation(operation string, old int) int {
	expression, _ := govaluate.NewEvaluableExpression(operation)
	parameters := make(map[string]any)
	parameters["old"] = old
	result, _ := expression.Evaluate(parameters)
	return int(result.(float64))
}

func getLevelOfMonkeyBusiness(monkey []*Monkey) int {
	inspections := functional.Map(monkey, func(monkey *Monkey) int {
		return monkey.Inspections
	})

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	return inspections[0] * inspections[1]
}
