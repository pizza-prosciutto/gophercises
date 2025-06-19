package quiz

import (
	"fmt"
	"io"
	"log"
)

func StartQuiz(problems []Problem, stdin io.Reader) Result {
	result := Result{0, 0}
	for idx, problem := range problems {
		result.TotalQuestions = result.TotalQuestions + 1
		fmt.Printf("Problem #%d: %s = ", idx, problem.Question)
		var answer string
		if _, err := fmt.Fscanln(stdin, &answer); err != nil {
			log.Fatalf("could not read input: %v", err)
			continue
		}
		if answer == problem.Answer {
			result.CorrectAnswers = result.CorrectAnswers + 1
		}
	}
	return result
}
