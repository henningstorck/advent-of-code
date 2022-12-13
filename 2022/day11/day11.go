package day11

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/henningstorck/advent-of-code/common/inputreader"
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

func Solve(reader *inputreader.InputReader) (int, int) {
	chunks := reader.ReadChunks(2022, 11)
	monkeys := Parse(chunks)
	return SolvePart1(monkeys), SolvePart2(monkeys)
}

func Parse(chunks [][]string) []*Monkey {
	return functional.Map(chunks, func(lines []string) *Monkey {
		monkey := &Monkey{}

		for _, line := range lines {
			if strings.HasPrefix(line, "Monkey ") {
				continue
			} else if strings.HasPrefix(line, "  Starting items: ") {
				startingItemsStr := line[18:]
				startingItemsStrSlice := strings.Split(startingItemsStr, ", ")
				startingItems := functional.Map(startingItemsStrSlice, func(valueStr string) int {
					value, _ := strconv.Atoi(valueStr)
					return value
				})
				monkey.Items = startingItems
			} else if strings.HasPrefix(line, "  Operation: new = ") {
				operation := line[19:]
				monkey.Operation = operation
			} else if strings.HasPrefix(line, "  Test: divisible by ") {
				divisibleByStr := line[21:]
				divisibleBy, _ := strconv.Atoi(divisibleByStr)
				monkey.DivisibleBy = divisibleBy
			} else if strings.HasPrefix(line, "    If true: throw to monkey ") {
				ifTrueStr := line[29:]
				ifTrue, _ := strconv.Atoi(ifTrueStr)
				monkey.IfTrue = ifTrue
			} else if strings.HasPrefix(line, "    If false: throw to monkey ") {
				ifFalseStr := line[30:]
				ifFalse, _ := strconv.Atoi(ifFalseStr)
				monkey.IfFalse = ifFalse
			}
		}

		return monkey
	})
}

func SolvePart1(monkeys []*Monkey) int {
	var monkeysCopy []*Monkey
	copier.CopyWithOption(&monkeysCopy, &monkeys, copier.Option{DeepCopy: true})

	return simulate(monkeysCopy, 20, func(item int) int {
		return item / 3
	})
}

func SolvePart2(monkeys []*Monkey) int {
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
