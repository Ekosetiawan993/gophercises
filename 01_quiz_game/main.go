package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// declare flag
	var csvFileName = flag.String("csv_name", "problems.csv", "type the csv file's name")
	var timeLimit = flag.Int("time_limit", 30, "type the maximal time limit")
	flag.Parse()

	// var userAnswer int
	var userScore int
	// chanellf or flag if the time is done
	timeUp := make(chan bool)

	// call the quiz engine function
	go func() {
		quizEngine(*csvFileName, &userScore)
		fmt.Printf("User final score: %v", userScore)
		os.Exit(0)
	}()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	select {
	case <-timer.C:
		// this C will block aftet 30
		fmt.Print("\nTime is Up.\n")
		fmt.Printf("User final score: %v", userScore)
		os.Exit(0)
	case <-timeUp:
		timer.Stop()
	}

	// fmt.Printf("User final score: %v", userScore)
}

func quizEngine(csvFileName string, userScore *int) {
	// open csv file
	f, err := os.Open(csvFileName)
	if err != nil {
		log.Fatal(err)
	}

	// close f in end of program
	defer f.Close()

	// read csv values
	csvReader := csv.NewReader(f)
	// fmt.Printf("type f %T", csvReader)

	for {
		// csv Read return string slice
		recOneLine, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// the question and real answer
		fmt.Printf("%v = ", recOneLine[0])
		realAnswer, err := strconv.Atoi(recOneLine[1])
		if err != nil {
			log.Fatal(err)
		}

		// scan the answer
		var userAnswer int
		fmt.Scan(&userAnswer)

		// check answer
		if userAnswer == realAnswer {
			*userScore++
		}

		clearInputBuffer()
	}

}

// function for clearing input buffer
func clearInputBuffer() {
	for {
		var temp rune
		_, err := fmt.Scanf("%c", &temp)
		if err != nil || temp == '\n' {
			break
		}
	}
}
