package dayxx_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/henningstorck/advent-of-code/downloader/dayxx"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	t.Skip()

	tests := []struct {
		filename  string
		expected1 int
		expected2 int
	}{
		{"example.txt", 0, 0},
		{"input.txt", 0, 0},
	}

	for _, test := range tests {
		reader := inputreader.NewInputReader("../..")
		result1, result2 := dayxx.Solve(reader, test.filename)
		assert.Equal(t, test.expected1, result1)
		assert.Equal(t, test.expected2, result2)
	}
}
