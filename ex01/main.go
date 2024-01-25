package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFilePath := flag.String("csv", "problems.csv", "a csv file in the format of 'problem'")
	flag.Parse()

	file := readFile(*csvFilePath)
	defer file.Close()
	problems := extractProblems(file)

	correct, incorrect := takeQuiz(problems)
	total := len(problems)
	printResults(total, correct, incorrect)
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(path string) *os.File {
	f, fileError := os.Open(path)
	errorChecker(fileError)
	return f
}

func extractProblems(file *os.File) []problem {
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	errorChecker(err)

	problems := make([]problem, len(data))
	for i, row := range data {
		problems[i] = problem{
			question: row[0],
			answer:   strings.TrimSpace(row[1]),
		}
	}
	return problems
}

func takeQuiz(problems []problem) (int, int) {
	correct := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		var ans string
		_, err := fmt.Scan(&ans)
		errorChecker(err)

		if ans == problem.answer {
			correct++
		}
	}
	incorrect := len(problems) - correct
	return correct, incorrect
}

func printResults(total int, correct int, incorrect int) {
	fmt.Println("\n<< RESULT >>")
	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Correct: %d\n", correct)
	fmt.Printf("Incorrect: %d\n", incorrect)
}
