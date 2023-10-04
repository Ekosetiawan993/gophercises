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

	// Open csv file
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Close f at the end of the program
	defer f.Close()

	// Read csv values
	csvReader := csv.NewReader(f)

	for {
		recOneLine, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v = ", recOneLine[0])

		fmt.Scan("input answer: %d ", &userAnswer)

		// Clear the input buffer
		clearInputBuffer()

		fmt.Printf("user input: %v \n", userAnswer)
	}
}

func clearInputBuffer() {
	// Read and discard any pending input until a newline character is encountered
	for {
		var temp rune
		_, err := fmt.Scanf("%c", &temp)
		if err != nil || temp == '\n' {
			break
		}
	}
}
