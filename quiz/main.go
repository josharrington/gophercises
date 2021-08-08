package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true
	lines, err := reader.ReadAll()
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

func timer(seconds int, correct *int, incorrect *int) {
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println("Time's up!")
	fmt.Println(fmt.Sprintf("\nYour score: %d/%d", *correct, *correct+*incorrect))
	os.Exit(0)
}

func main() {
	timeout := flag.Int("timeout", -1, "The number of seconds to run the quiz game. Disabled by default.")
	shuffle := flag.Bool("shuffle", false, "Shuffle the questions.")
	input := flag.String("input", "./problems.csv", "The path to the quiz questions csv.")
	flag.Parse()

	questions, err := parseCsv(*input)
	if err != nil {
		fmt.Println("Error parsing contents of csv: ", err)
		os.Exit(1)
	}

	fmt.Println("Welcome to the quiz! You'll be given questions you have to answer. Your score will be displayed at the end.")
	fmt.Print("Press ENTER when you're ready to begin!")
	fmt.Scanln()

	var answer string
	var correct int
	total := len(questions)

	if *shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })
	}

	if *timeout > 0 {
		go timer(*timeout, &correct, &total)
	}

	for _, v := range questions {
		fmt.Print(fmt.Sprintf("%s: ", v.problem))
		fmt.Scanln(&answer)
		answer = strings.TrimSpace(answer)

		if strings.ToLower(answer) == strings.ToLower(v.answer) {
			correct += 1
			fmt.Println("Correct!")
		} else {
			fmt.Println("Wrong :(")
		}
	}

	fmt.Println(fmt.Sprintf("\nYour score: %d/%d", correct, correct+total))
}
