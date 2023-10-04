package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type csvContent struct {
	question   string
	realAnswer int
}

func main() {
	var userAnswer int
	// var csvContents []csvContent

	// open csv file
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// close f in end of program
	defer f.Close()

	// read csv values
	csvReader := csv.NewReader(f)
	for {
		// csv Read return string slice
		rec_one_line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v = ", rec_one_line[0])
		fmt.Scanf("%d", &userAnswer)
		fmt.Printf("user input: %v \n", userAnswer)
	}
}
