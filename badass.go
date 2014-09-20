package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	SATURATION  = 72
	WEEK        = 7
	COLS        = 75
	MAX_COMMITS = 74
	START_DATE  = ""
)

var (
	flSaturation int
	flColumns    int
	flMaxCommits int
	flStartDate  string
	totalCommits int
)

const (
	UsgSaturation      = "sets saturation of working days range 1-100"
	UsgSaturationShort = "short hand version of --saturation"
	UsgColumns         = "sets the value of maximum rows for contribution graph"
	UsgColumnsShort    = "short hand version of --columns"
	UsgStartDate       = "sets the starting date of contribution graph. Must be in MM/DD/YYYY format"
	UsgStartDateShort  = "short hand version of --date"
	UsgMaxCommits      = "sets the value of maximum commits for a date"
	UsgMaxCommitsShort = "short hand version of --max-commits"
)

func init() {
	flag.IntVar(&flSaturation, "saturation", SATURATION, UsgSaturation)
	flag.IntVar(&flSaturation, "s", SATURATION, UsgSaturationShort)

	flag.IntVar(&flColumns, "columns", COLS, UsgColumns)
	flag.IntVar(&flColumns, "c", COLS, UsgColumnsShort)

	flag.StringVar(&flStartDate, "date", START_DATE, UsgStartDate)
	flag.StringVar(&flStartDate, "d", START_DATE, UsgStartDateShort)

	flag.IntVar(&flMaxCommits, "max-commits", MAX_COMMITS, UsgMaxCommits)
	flag.IntVar(&flMaxCommits, "m", MAX_COMMITS, UsgMaxCommitsShort)

	flag.Parse()

	if flSaturation < 1 || flSaturation > 100 {
		flSaturation = SATURATION
	}
	_, err := time.Parse("01/02/2006", flStartDate)

	if flStartDate == "" {
		t := time.Now()
		flStartDate = fmt.Sprintf("%02d/%02d/%04d", t.Month(), t.Day(), t.Year()-1)
	} else if err != nil {
		fmt.Println("please provide a valid date in MM/DD/YYYY format.")
		os.Exit(2)

	}
}

func main() {
	beginTime := time.Now()
	totalCommits = 0

	rand.Seed(beginTime.Unix())

	var pattern string

	for i := 0; i < flColumns; i++ {
		for j := 0; j < WEEK; j++ {
			if rand.Intn(100) > flSaturation {
				pattern += "."
			} else {
				pattern += "0"
			}
		}
		pattern += "\n"
	}

	weeks := strings.Split(pattern, "\n")
	date := parseTime(flStartDate)

	for i := 0; i < len(weeks)-1; i++ {
		week := weeks[i]

		for j := range week {
			cell := (string)(week[j])

			if cell == "0" {
				doCommits(date)
			}
			date = nextDay(date)
		}
	}

	endTime := time.Now()
	fmt.Printf("\nBoom! %d commits generated in %s duration.\n", totalCommits, (endTime.Sub(beginTime)).String())

}
func doCommits(date time.Time) {
	date = addRandomDuration(date)
	r := rand.Intn(flMaxCommits)
	for i := 0; i < r; i++ {
		formatted := formatTime(date)

		os.Setenv("GIT_AUTHOR_DATE", formatted)
		os.Setenv("GIT_COMMITTER_DATE", formatted)

		msg := fmt.Sprintf("Commit: %02d/%02d (max: %02d) for Date: %s", i, r, flMaxCommits, date.Format(time.RFC822))
		cmd := exec.Command("git", "commit", "--allow-empty", "-m "+msg)

		out, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}

		fmt.Print(string(out))
		totalCommits++

		date = addRandomDuration(date)
	}

}

func parseTime(value string) time.Time {
	val, err := time.Parse("01/02/2006", value)
	if err != nil {
		panic(err)
	}
	return val
}

func formatTime(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

}

func nextDay(value time.Time) time.Time {
	return time.Date(value.Year(), value.Month(), value.Day()+1, 0, 0, 0, 0, time.UTC)
}

func addRandomDuration(value time.Time) time.Time {
	t := time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, time.UTC)

	h := rand.Intn(24)
	m := rand.Intn(60)
	s := rand.Intn(60)

	d := parseDuration(h, m, s)

	return t.Add(d)
}

func parseDuration(h int, m int, s int) time.Duration {
	d := fmt.Sprintf("%dh%dm%ds", h, m, s)
	duration, err := time.ParseDuration(d)
	if err != nil {
		panic(err)
	}

	return duration
}
