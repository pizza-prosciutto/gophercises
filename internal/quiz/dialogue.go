package quiz

import (
	"fmt"
	"os"
	"strings"
	"time"
)

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
