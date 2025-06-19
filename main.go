package main

import (
	"fmt"
	"os"

	"github.com/pizza-prosciutto/gophercises/internal/quiz"
)

func main() {
	problems, err := quiz.ReadProblems("assets/quiz/problems.csv")
	if err != nil {
		os.Exit(1)
	}
	result := quiz.StartQuiz(problems, os.Stdin)
	fmt.Printf("You scored %d out of %d.", result.CorrectAnswers, result.TotalQuestions)
}
