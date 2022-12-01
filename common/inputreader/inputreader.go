package inputreader

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
)

const filenameTemplate = "%d/day%02d/input.txt"

type InputReader struct {
	FS fs.FS
}

func NewInputReader(dir string) *InputReader {
	return &InputReader{os.DirFS(dir)}
}

func (reader *InputReader) ReadLines(year, day int) []string {
	return read(reader.FS, getFilename(year, day), bufio.ScanLines)
}

func (reader *InputReader) ReadWords(year, day int) []string {
	return read(reader.FS, getFilename(year, day), bufio.ScanWords)
}

func (reader *InputReader) ReadRunes(year, day int) []string {
	return read(reader.FS, getFilename(year, day), bufio.ScanRunes)
}

func (reader *InputReader) ReadChunks(year, day int) [][]string {
	lines := reader.ReadLines(year, day)
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

func getFilename(year, day int) string {
	return fmt.Sprintf(filenameTemplate, year, day)
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
