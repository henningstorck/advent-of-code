package day13

import (
	"encoding/json"

	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/henningstorck/advent-of-code/common/sorting"
)

func Solve(reader *inputreader.InputReader) (int, int) {
	chunks := reader.ReadChunks(2022, 13)
	return SolvePart1(chunks), SolvePart2(chunks)
}

func SolvePart1(chunks [][]string) int {
	indices := []int{}

	for i, pair := range chunks {
		left := unmarshalPacket(pair[0])
		right := unmarshalPacket(pair[1])

		if compare(left, right) < 0 {
			indices = append(indices, i+1)
		}
	}

	return functional.Reduce(indices, func(sum, i int) int {
		return sum + i
	}, 0)
}

func SolvePart2(chunks [][]string) int {
	packets := functional.Reduce(chunks, func(packets []any, pair []string) []any {
		left := unmarshalPacket(pair[0])
		right := unmarshalPacket(pair[1])
		return append(packets, left, right)
	}, []any{})

	packets = append(packets, unmarshalPacket("[[2]]"))
	packets = append(packets, unmarshalPacket("[[6]]"))

	packets = sorting.BubbleSort(packets, func(a, b any) bool {
		return compare(a, b) > 0
	})

	packetStrs := functional.Map(packets, marshalPacket)

	i := indexOf(packetStrs, "[[2]]") + 1
	j := indexOf(packetStrs, "[[6]]") + 1

	return i * j
}

func compare(left any, right any) int {
	leftInt, leftIntOk := left.(float64)
	rightInt, rightIntOk := right.(float64)

	if leftIntOk && rightIntOk {
		return int(leftInt - rightInt)
	}

	leftSlice, leftSliceOk := left.([]any)
	rightSlice, rightSliceOk := right.([]any)

	if leftSliceOk && rightSliceOk {
		if len(leftSlice) == 0 && len(rightSlice) == 0 {
			return 0
		}

		if len(leftSlice) == 0 {
			return -1
		}

		if len(rightSlice) == 0 {
			return 1
		}

		result := compare(leftSlice[0], rightSlice[0])

		if result == 0 {
			return compare(leftSlice[1:], rightSlice[1:])
		}

		return result
	}

	if leftSliceOk {
		return compare(left, []any{right})
	}

	if rightSliceOk {
		return compare([]any{left}, right)
	}

	return 0
}

func indexOf(packets []string, a string) int {
	for i, b := range packets {
		if a == b {
			return i
		}
	}

	return -1
}

func unmarshalPacket(packetStr string) any {
	var packet any
	json.Unmarshal([]byte(packetStr), &packet)
	return packet
}

func marshalPacket(packet any) string {
	packetStr, _ := json.Marshal(packet)
	return string(packetStr)
}
