package dayxx_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/henningstorck/advent-of-code/downloader/dayxx"
	"github.com/stretchr/testify/assert"
)

func TestSolveParts(t *testing.T) {
	lines := []string{
		"",
	}

	resultPart1 := dayxx.SolvePart1(lines)
	resultPart2 := dayxx.SolvePart2(lines)
	assert.Equal(t, 0, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	t.Skip()
	reader := inputreader.NewInputReader("../..")
	resultPart1, resultPart2 := dayxx.Solve(reader)
	assert.Equal(t, 0, resultPart1)
	assert.Equal(t, 0, resultPart2)
}
