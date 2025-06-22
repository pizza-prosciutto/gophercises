package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pizza-prosciutto/gophercises/internal/quiz"
)

func main() {

	var timeout time.Duration
	var csvPath string

	flag.DurationVar(&timeout, "timeout", 10*time.Second, "quiz timeout, e.g. 10s, 1m")
	flag.StringVar(&csvPath, "csv", "problems.csv", "path to quiz questions in csv format")
	flag.Parse()

	result, err := quiz.StartQuiz(timeout, csvPath)
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("You scored %d out of %d.\n", result.CorrectAnswers, result.TotalQuestions)
}
