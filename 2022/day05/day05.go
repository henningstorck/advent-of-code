package day05

import (
	"bytes"
	"regexp"
	"strconv"

	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/henningstorck/advent-of-code/common/stack"
	"github.com/jinzhu/copier"
)

type Instruction struct {
	Amount int
	From   int
	To     int
}

func Solve(reader *inputreader.InputReader, filename string) (string, string) {
	chunks := reader.ReadChunks(2022, 5, filename)
	stacks, instructions := parse(chunks)
	return solvePart1(stacks, instructions), solvePart2(stacks, instructions)
}

func parse(chunks [][]string) ([][]string, []Instruction) {
	piles := [][]string{}
	instructions := []Instruction{}

	for pos := 1; pos < len(chunks[0][0]); pos += 4 {
		piles = append(piles, parsePile(chunks[0], pos))
	}

	for _, chunk := range chunks[1] {
		instructions = append(instructions, parseInstruction(chunk))
	}

	return piles, instructions
}

func parsePile(lines []string, pos int) []string {
	pile := stack.Stack[string]{}

	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		value := line[pos : pos+1]

		if value != " " {
			pile = pile.Prepend(line[pos : pos+1])
		}
	}

	return pile
}

func parseInstruction(line string) Instruction {
	regex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	results := regex.FindStringSubmatch(line)
	amount, _ := strconv.Atoi(results[1])
	from, _ := strconv.Atoi(results[2])
	to, _ := strconv.Atoi(results[3])
	return Instruction{amount, from, to}
}

func solvePart1(piles [][]string, instructions []Instruction) string {
	var pilesCopy []stack.Stack[string]
	copier.CopyWithOption(&pilesCopy, &piles, copier.Option{DeepCopy: true})

	for _, instruction := range instructions {
		from := instruction.From - 1
		to := instruction.To - 1

		for i := 0; i < instruction.Amount; i++ {
			pileFrom := pilesCopy[from]
			pileFrom, value, _ := pileFrom.Pop()
			pilesCopy[from] = pileFrom
			pileTo := pilesCopy[to]
			pileTo = pileTo.Push(value)
			pilesCopy[to] = pileTo
		}
	}

	return pilesToStrings(pilesCopy)
}

func solvePart2(piles [][]string, instructions []Instruction) string {
	var pilesCopy []stack.Stack[string]
	copier.CopyWithOption(&pilesCopy, &piles, copier.Option{DeepCopy: true})

	for _, instruction := range instructions {
		helperStack := stack.Stack[string]{}
		from := instruction.From - 1
		to := instruction.To - 1

		for i := 0; i < instruction.Amount; i++ {
			pileFrom := pilesCopy[from]
			pileFrom, value, _ := pileFrom.Pop()
			pilesCopy[from] = pileFrom
			helperStack = helperStack.Push(value)
		}

		for i := 0; i < instruction.Amount; i++ {
			var value string
			helperStack, value, _ = helperStack.Pop()
			pileTo := pilesCopy[to]
			pileTo = pileTo.Push(value)
			pilesCopy[to] = pileTo
		}
	}

	return pilesToStrings(pilesCopy)
}

func pilesToStrings(piles []stack.Stack[string]) string {
	var buffer bytes.Buffer

	for _, pile := range piles {
		value, _ := pile.Peek()
		buffer.WriteString(value)
	}

	return buffer.String()
}
