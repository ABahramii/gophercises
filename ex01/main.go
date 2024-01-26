package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

type result struct {
	total     int
	correct   int
	incorrect int
}

func main() {
	csvFilePath := flag.String("csv", "problems.csv", "a csv file in the format of 'problem'")
	timeLimit := flag.Int("limit", 30, "time limit for quiz")
	flag.Parse()

	file := readFile(*csvFilePath)
	defer file.Close()
	problems := extractProblems(file)

	resChan := make(chan result)
	go takeQuiz(problems, resChan)

	select {
	case res := <-resChan:
		printResults(res)
	case <-time.After(time.Duration(*timeLimit) * time.Second):
		fmt.Println("\n\nQuiz time has expired :(")
	}
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
	data := readCsvFile(file)
	problems := make([]problem, len(data))
	for i, row := range data {
		problems[i] = problem{
			question: row[0],
			answer:   strings.TrimSpace(row[1]),
		}
	}
	return problems
}

func readCsvFile(file *os.File) [][]string {
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	errorChecker(err)
	return data
}

func takeQuiz(problems []problem, res chan result) {
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

	total := len(problems)
	incorrect := total - correct
	res <- result{
		total:     total,
		correct:   correct,
		incorrect: incorrect,
	}
}

func printResults(res result) {
	fmt.Println("\n<< RESULT >>")
	fmt.Printf("Total: %d\n", res.total)
	fmt.Printf("Correct: %d\n", res.correct)
	fmt.Printf("Incorrect: %d\n", res.incorrect)
}
