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

func Solve(reader *inputreader.InputReader) (string, string) {
	chunks := reader.ReadChunks(2022, 5)
	stacks, instructions := Parse(chunks)
	return SolvePart1(stacks, instructions), SolvePart2(stacks, instructions)
}

func Parse(chunks [][]string) ([]*stack.Stack, []Instruction) {
	stacks := []*stack.Stack{}
	instructions := []Instruction{}

	for pos := 1; pos < len(chunks[0][0]); pos += 4 {
		stacks = append(stacks, parseStack(chunks[0], pos))
	}

	for _, chunk := range chunks[1] {
		instructions = append(instructions, parseInstruction(chunk))
	}

	return stacks, instructions
}

func parseStack(lines []string, pos int) *stack.Stack {
	stack := &stack.Stack{}

	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		value := line[pos : pos+1]

		if value != " " {
			stack.Prepend(line[pos : pos+1])
		}
	}

	return stack
}

func parseInstruction(line string) Instruction {
	regex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	results := regex.FindStringSubmatch(line)
	amount, _ := strconv.Atoi(results[1])
	from, _ := strconv.Atoi(results[2])
	to, _ := strconv.Atoi(results[3])
	return Instruction{amount, from, to}
}

func SolvePart1(stacks []*stack.Stack, instructions []Instruction) string {
	var stacksCopy []*stack.Stack
	copier.CopyWithOption(&stacksCopy, &stacks, copier.Option{DeepCopy: true})

	for _, instruction := range instructions {
		from := instruction.From - 1
		to := instruction.To - 1

		for i := 0; i < instruction.Amount; i++ {
			value, _ := stacksCopy[from].Pop()
			stacksCopy[to].Push(value)
		}
	}

	return stacksToStrings(stacksCopy)
}

func SolvePart2(stacks []*stack.Stack, instructions []Instruction) string {
	var stacksCopy []*stack.Stack
	copier.CopyWithOption(&stacksCopy, &stacks, copier.Option{DeepCopy: true})

	for _, instruction := range instructions {
		helperStack := &stack.Stack{}
		from := instruction.From - 1
		to := instruction.To - 1

		for i := 0; i < instruction.Amount; i++ {
			value, _ := stacksCopy[from].Pop()
			helperStack.Push(value)
		}

		for i := 0; i < instruction.Amount; i++ {
			value, _ := helperStack.Pop()
			stacksCopy[to].Push(value)
		}
	}

	return stacksToStrings(stacksCopy)
}

func stacksToStrings(stacks []*stack.Stack) string {
	var buffer bytes.Buffer

	for _, stack := range stacks {
		value, _ := stack.Peek()
		buffer.WriteString(value)
	}

	return buffer.String()
}
