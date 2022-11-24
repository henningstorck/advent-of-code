package main

import (
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
	Day     int
	Year    int
	Session string
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
