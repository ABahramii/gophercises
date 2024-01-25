package main

import (
	"bufio"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func main() {
	readFile, err := os.Open("./ex_01/problems.csv")
	errorChecker(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	problems := make(map[int]problem)

	i := 1
	for fileScanner.Scan() {
		split := strings.Split(fileScanner.Text(), ",")
		p := problem{question: split[0], answer: split[1]}
		problems[i] = p
		i++
	}

	/*correct, incorrect := 0, 0
	for num, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", num, problem.question)
		var ans string
		_, err := fmt.Scan(&ans)
		errorChecker(err)

		if ans == problem.answer {
			correct++
		} else {
			incorrect++
		}
	}

	fmt.Printf("Total: %d\n", len(problems))
	fmt.Printf("Correct: %d\n", correct)
	fmt.Printf("Incorrect: %d\n", incorrect)*/
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
