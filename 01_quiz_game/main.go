package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type csvContent struct {
	question   string
	realAnswer int
}

func main() {
	// declare flag
	var csvFileName = flag.String("csv_name", "problems.csv", "type the csv file's name")
	var timeLimit = flag.Int("time_limit", 30, "type the maximal time limit")
	var shuffleQuestion = flag.Bool("shuffle_question", true, "shuffling question order or not")
	flag.Parse()

	// var userAnswer int
	var userScore int
	// chanellf or flag if the time is done
	timeUp := make(chan bool)

	// fmt.Printf("%T %v", shuffleQuestion, *shuffleQuestion)
	fmt.Printf("You have %v seconds. Press enter to start the quiz.....", *timeLimit)
	fmt.Scanln()

	f, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if !*shuffleQuestion {
		// call the quiz engine function
		go func() {
			quizEngine(f, &userScore)
			fmt.Printf("User final score: %v", userScore)
			os.Exit(0)
		}()
	} else {
		go func() {
			quizEngineShuffle(f, &userScore)
			fmt.Printf("User final score: %v", userScore)
			os.Exit(0)
		}()
	}

	// go func() {
	// 	quizEngineShuffle(*csvFileName, &userScore)
	// 	fmt.Printf("User final score: %v", userScore)
	// 	os.Exit(0)
	// }()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	select {
	case <-timer.C:
		// this C will block aftet 30
		fmt.Print("\nTime is Up.\n")
		fmt.Printf("User final score: %v", userScore)
		defer fmt.Println("Close f")
		defer f.Close()
		os.Exit(0)
	case <-timeUp:
		timer.Stop()
	}

	// fmt.Printf("User final score: %v", userScore)
}

func quizEngineShuffle(f *os.File, userScore *int) {
	var questionsList []csvContent

	// open csv file
	// f, err := os.Open(csvFileName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// close f in end of program
	// defer fmt.Println("close f in engine")

	// read csv values
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("type f %T", csvReader)

	for _, line := range data {
		var rec csvContent

		for j, field := range line {
			if j == 0 {
				rec.question = field
			} else if j == 1 {
				rec.realAnswer, err = strconv.Atoi(field)
				if err != nil {
					log.Fatal(err)
				}
			}

		}
		questionsList = append(questionsList, rec)
	}

	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	randomizer.Shuffle(len(questionsList), func(i, j int) {
		questionsList[i], questionsList[j] = questionsList[j], questionsList[i]
	})

	for _, question := range questionsList {
		// fmt.Println(question.question)
		// the question and real answer
		fmt.Printf("%v = ", question.question)
		realAnswer := question.realAnswer

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

func quizEngine(f *os.File, userScore *int) {
	// open csv file

	// fmt.Printf("type of f %T", f)

	// close f in end of program
	// defer fmt.Println("close f in engine")
	// defer f.Close()

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
		var userAnswer string
		fmt.Scan(&userAnswer)
		// fmt.Printf("user input %v, %T", userAnswer, userAnswer)

		userAnswerInt, err := strconv.Atoi(userAnswer)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// check answer
		if userAnswerInt == realAnswer {
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
