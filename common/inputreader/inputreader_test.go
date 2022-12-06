package inputreader_test

import (
	"testing"
	"testing/fstest"

	"github.com/henningstorck/advent-of-code/common/inputreader"
	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	fs := fstest.MapFS{
		"2022/day01/input.txt": {
			Data: []byte("hello world\nhello universe\n\nhello everything"),
		},
	}

	reader := &inputreader.InputReader{fs}
	lines := reader.ReadLines(2022, 1)
	assert.Equal(t, []string{"hello world", "hello universe", "", "hello everything"}, lines)
}

func TestReadWords(t *testing.T) {
	fs := fstest.MapFS{
		"2022/day01/input.txt": {
			Data: []byte("hello world\nhello universe\n\nhello everything"),
		},
	}

	reader := &inputreader.InputReader{fs}
	words := reader.ReadWords(2022, 1)
	assert.Equal(t, []string{"hello", "world", "hello", "universe", "hello", "everything"}, words)
}

func TestReadRunes(t *testing.T) {
	fs := fstest.MapFS{
		"2022/day01/input.txt": {
			Data: []byte("hello\nworld"),
		},
	}

	reader := &inputreader.InputReader{fs}
	runes := reader.ReadRunes(2022, 1)
	assert.Equal(t, []rune{'h', 'e', 'l', 'l', 'o', '\n', 'w', 'o', 'r', 'l', 'd'}, runes)
}

func TestReadChunks(t *testing.T) {
	fs := fstest.MapFS{
		"2022/day01/input.txt": {
			Data: []byte("hello world\nhello universe\n\nhello everything"),
		},
	}

	reader := &inputreader.InputReader{fs}
	chunks := reader.ReadChunks(2022, 1)
	assert.Len(t, chunks, 2)
	assert.Equal(t, []string{"hello world", "hello universe"}, chunks[0])
	assert.Equal(t, []string{"hello everything"}, chunks[1])
}
