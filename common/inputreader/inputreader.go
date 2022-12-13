package inputreader

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/henningstorck/advent-of-code/common/functional"
)

const filenameTemplate = "%d/day%02d/%s"

type InputReader struct {
	FS fs.FS
}

func NewInputReader(dir string) *InputReader {
	return &InputReader{os.DirFS(dir)}
}

func (reader *InputReader) ReadLines(year, day int, filename string) []string {
	return read(reader.FS, getFilename(year, day, filename), bufio.ScanLines)
}

func (reader *InputReader) ReadWords(year, day int, filename string) []string {
	return read(reader.FS, getFilename(year, day, filename), bufio.ScanWords)
}

func (reader *InputReader) ReadRunes(year, day int, filename string) []rune {
	chars := read(reader.FS, getFilename(year, day, filename), bufio.ScanRunes)

	return functional.Map(chars, func(char string) rune {
		return rune(char[0])
	})
}

func (reader *InputReader) ReadChunks(year, day int, filename string) [][]string {
	lines := reader.ReadLines(year, day, filename)
	chunk := []string{}
	chunks := [][]string{}

	for _, line := range lines {
		if len(line) == 0 {
			chunks = append(chunks, chunk)
			chunk = []string{}
		} else {
			chunk = append(chunk, line)
		}
	}

	chunks = append(chunks, chunk)
	return chunks

}

func getFilename(year, day int, filename string) string {
	return fmt.Sprintf(filenameTemplate, year, day, filename)
}

func read(fs fs.FS, filename string, splitFunc bufio.SplitFunc) []string {
	file, err := fs.Open(filename)

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(splitFunc)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return lines
}
