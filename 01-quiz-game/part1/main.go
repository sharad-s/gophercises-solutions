package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	// Create empty slice with determined length
	ret := make([]problem, len(lines))
	// Populate lines into slices
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	// Return slice
	return ret
}

func main() {
	// Create and Parse through flags
	csvFilename := flag.String("f", "problems.csv", "a csv in the format of 'question, answer'")
	flag.Parse()

	// Open CSV file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	// Parse CSV File
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)

	correct := 0
	// Ask each question from problems
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	// Print result
	fmt.Printf("You got %d out of %d questions right.\n", correct, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
