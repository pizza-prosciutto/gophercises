// Solution for the Quiz Game
package quiz

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Problem represents a single quiz question and its answer.
type Problem struct {
	Question string
	Answer   string
}

// Result stores the quiz score.
type Result struct {
	TotalQuestions int
	CorrectAnswers int
}

// readProblems reads a CSV file and maps the records to a slice of Problems.
func readProblems(csvPath string) ([]Problem, error) {
	if _, err := os.Stat(csvPath); err != nil {
		return nil, errors.New("provided csv file does not exist")
	}
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return nil, fmt.Errorf("could not open csv file: %v", err)
	}
	var problems []Problem
	csvReader := csv.NewReader(csvFile)
	for {
		csvRecord, err := csvReader.Read()
		if err == io.EOF {
			return problems, nil
		}
		if err != nil {
			fmt.Printf("Could not read record: %v\n", err)
		} else {
			problems = append(problems, Problem{Question: csvRecord[0], Answer: csvRecord[1]})
		}
	}
}

// showDialog prints the questions to STDOUT, and reads the answers from STDIN.
// The channel is closed when all the questions have been answered.
func showDialogue(problems []Problem, result *Result, done chan any) {
	for idx, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", idx+1, problem.Question)
		var answer string
		if _, err := fmt.Fscanln(os.Stdin, &answer); err != nil {
			fmt.Printf("Could not read input: %v", err)
			continue
		}
		answer = strings.TrimSpace(answer)
		if strings.EqualFold(answer, problem.Answer) {
			(*result).CorrectAnswers++
		}
	}
	close(done)
}

// StartQuiz sets a timeout and starts the quiz. It returns the result
// when all the questions have been answered or timeout has expired.
func StartQuiz(timeout time.Duration, csvPath string) (Result, error) {
	problems, err := readProblems(csvPath)
	if err != nil {
		return Result{}, fmt.Errorf("failed to read problems: %v", err)
	}
	result := Result{TotalQuestions: len(problems)}
	timer := time.NewTimer(timeout)
	done := make(chan any)
	go showDialogue(problems, &result, done)
	for {
		select {
		case <-timer.C:
			fmt.Println("\nTime is up!")
			return result, nil
		case _, ok := <-done:
			if !ok {
				return result, nil
			}
		}
	}
}
