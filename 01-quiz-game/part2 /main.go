package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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
	timeLimit := flag.Int("t", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	// Open CSV file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	// Parse CSV File and generate problems
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)

	// Create timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	// Ask each question from problems
	for i, p := range problems {

		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

		// Create non-blocking Goroutine for scanning answer
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		// Run logic based on goroutines
		select {
		// If a message is received from the timer, end quiz.
		case <-timer.C:
			fmt.Printf("\nYou got %d out of %d questions right.\n", correct, len(problems))
			return

		case answer := <-answerCh:
			// Otherwise keep asking questions
			if answer == p.a {
				correct++
			}
		}
	}

	// Print result
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
