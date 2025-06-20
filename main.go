package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pizza-prosciutto/gophercises/internal/quiz"
)

func main() {
	result, err := quiz.StartQuiz(2*time.Second, "assets/quiz/problems.csv")
	if err != nil {
		fmt.Printf("Oops: %v", err)
		os.Exit(1)
	}
	fmt.Printf("You scored %d out of %d.\n", result.CorrectAnswers, result.TotalQuestions)
}
