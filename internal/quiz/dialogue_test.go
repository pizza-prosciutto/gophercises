package quiz_test

import (
	"strings"
	"testing"

	"github.com/pizza-prosciutto/gophercises/internal/quiz"
)

func TestStartQuiz(t *testing.T) {
	problems := []quiz.Problem{
		{Question: "1+1", Answer: "2"},
	}
	answer := strings.NewReader("2\n")
	result := quiz.StartQuiz(problems, answer)
	if result.CorrectAnswers != 1 {
		t.Errorf("expected %d correct answers, but got: %d", 1, result.CorrectAnswers)
	}
	if result.TotalQuestions != 1 {
		t.Errorf("expected %d total questions, but got: %d", 1, result.TotalQuestions)
	}
}
