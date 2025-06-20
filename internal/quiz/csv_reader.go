package quiz

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

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
