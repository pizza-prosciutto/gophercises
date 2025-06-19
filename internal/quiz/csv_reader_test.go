package quiz_test

import (
	"os"
	"testing"

	"github.com/pizza-prosciutto/gophercises/internal/quiz"
)

func TestReadProblems(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test-*.csv")
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	testData := "5+5,10\n7+3,10\n1+1,2"
	if _, err := tempFile.WriteString(testData); err != nil {
		t.Fatalf("failed to write to test file: %v", err)
	}

	if err := tempFile.Close(); err != nil {
		t.Fatalf("failed to close test file: %v", err)
	}

	problems, err := quiz.ReadProblems(tempFile.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(problems) != 3 {
		t.Errorf("expected %d records, but got %d", 3, len(problems))
	}

	if problems[1].Question != "7+3" {
		t.Errorf("expected question '%s', but got '%s'", "7+3", problems[1].Question)
	}

	if problems[1].Answer != "10" {
		t.Errorf("expected answer '%s', but got '%s'", "10", problems[1].Answer)
	}
}
