package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

const (
	inputUrl         = "https://adventofcode.com/%d/day/%d/input"
	filenameTemplate = "%d/day%02d/input.txt"
)

type Flags struct {
	Day       int
	Year      int
	Session   string
	Bootstrap bool
}

func main() {
	godotenv.Load()
	flags := getFlags()

	if flags.Day < 1 || flags.Day > 25 {
		log.Fatalln("day has to be between 0 and 25")
	}

	body, err := loadInput(flags.Day, flags.Year, flags.Session)

	if err != nil {
		log.Fatalln(err)
	}

	defer body.Close()
	err = writeFile(fmt.Sprintf(filenameTemplate, flags.Year, flags.Day), body)

	if err != nil {
		log.Fatalln(err)
	}

	if flags.Bootstrap {
		err = bootstrap(flags.Day, flags.Year)

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func getFlags() Flags {
	flags := Flags{}
	curYear, curMonth, curDay := time.Now().Date()
	defaultSession := os.Getenv("AOC_SESSION")

	if curMonth != 12 {
		curYear--
	}

	flag.IntVar(&flags.Day, "day", curDay, "Day to download")
	flag.IntVar(&flags.Year, "year", curYear, "Year to download")
	flag.StringVar(&flags.Session, "session", defaultSession, "AoC session")
	flag.BoolVar(&flags.Bootstrap, "bootstrap", true, "Bootstrap source files")
	flag.Parse()
	return flags
}

func loadInput(day, year int, session string) (io.ReadCloser, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(inputUrl, year, day), nil)

	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("got HTTP status code %d", res.StatusCode)
	}

	return res.Body, nil
}

func writeFile(filename string, body io.Reader) error {
	err := os.MkdirAll(filepath.Dir(filename), 0755)

	if err != nil {
		return err
	}

	out, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer out.Close()
	_, err = io.Copy(out, body)

	if err != nil {
		return err
	}

	return nil
}

func bootstrap(day, year int) error {
	err := copyTemplate(day, year, "downloader/dayxx/dayxx.go", fmt.Sprintf("%d/day%02d/day%02d.go", year, day, day))

	if err != nil {
		return err
	}

	err = copyTemplate(day, year, "downloader/dayxx/dayxx_test.go", fmt.Sprintf("%d/day%02d/day%02d_test.go", year, day, day))

	if err != nil {
		return err
	}

	return nil
}

func copyTemplate(day, year int, src, dest string) error {
	contents, err := os.ReadFile(src)

	if err != nil {
		return err
	}

	contents = bytes.Replace(contents, []byte("downloader/dayxx"), []byte(fmt.Sprintf("%d/day%02d", year, day)), -1)
	contents = bytes.Replace(contents, []byte("dayxx"), []byte(fmt.Sprintf("day%02d", day)), -1)
	contents = bytes.Replace(contents, []byte("reader.ReadLines(0, 0)"), []byte(fmt.Sprintf("reader.ReadLines(%d, %d)", year, day)), -1)
	contents = bytes.Replace(contents, []byte("\tt.Skip()\n"), []byte(""), -1)
	err = os.WriteFile(dest, contents, 0644)

	if err != nil {
		return err
	}

	return nil
}
