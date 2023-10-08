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

func main() {
	csvFilename := flag.String("csv_name", "problems.csv", "type csv filename")
	timeLimit := flag.Int("time_limit", 30, "time limit in seconds")
	shuffleP := flag.Bool("shuffle", false, "shuffle teh questions order")
	flag.Parse()

	f, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintln("Failed to open csv: %s\n", *csvFilename))
	}

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse csv file")
	}

	problems := parseLines(lines)

	if *shuffleP {
		problems = shuffleProblems(problems)
	}

	fmt.Printf("You have %d seconds to finish it, press enter to start the timer", *timeLimit)
	fmt.Scanln()

	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime is up.")
			fmt.Printf("Your scores %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}

	}

	fmt.Printf("Your scores %d out of %d.\n", correct, len(problems))

}

func shuffleProblems(problems []problem) []problem {
	shuffledProblems := problems
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	randomizer.Shuffle(len(shuffledProblems), func(i, j int) {
		shuffledProblems[i], shuffledProblems[j] = shuffledProblems[j], shuffledProblems[i]
	})

	return shuffledProblems
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
