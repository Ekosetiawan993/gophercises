package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

type csvContent struct {
	question   string
	realAnswer int
}

func main() {
	// var userAnswer int
	var userScore int
	// alarm to wait until the quizEngine done
	var alarm sync.WaitGroup

	// default duartion
	// durationOfTime := time.Duration(30) * time.Second

	// call the quiz engine function
	alarm.Add(1)
	go quizEngine("problems.csv", &userScore, &alarm)
	alarm.Wait()

	fmt.Printf("User final score: %v", userScore)
}

func quizEngine(csvFileName string, userScore *int, alarm *sync.WaitGroup) {
	// open csv file
	f, err := os.Open("problems.csv")
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

	alarm.Done()
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
