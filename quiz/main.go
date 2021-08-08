package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Question struct {
	problem string
	answer  string
}

func parseCsv(filename string) ([]Question, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	var questions []Question

	for _, v := range lines {
		question := Question{problem: v[0], answer: v[1]}
		questions = append(questions, question)
	}

	return questions, nil
}

func main() {
	result, err := parseCsv("./problems.csv")
	if err != nil {
		fmt.Println("Error parsing contents of csv: ", err)
		os.Exit(1)
	}

	fmt.Println("Welcome to the quiz! You'll be given questions you have to answer. Your score will be displayed at the end.")

	var answer string
	var correct int
	var incorrect int

	for _, v := range result {
		fmt.Print(fmt.Sprintf("%s: ", v.problem))
		fmt.Scanln(&answer)
		if answer == v.answer {
			correct += 1
			fmt.Println("Correct!")
		} else {
			incorrect += 1
			fmt.Println("Wrong :(")
		}
	}

	fmt.Println(fmt.Sprintf("\nYour score: %d/%d", correct, correct+incorrect))
}
