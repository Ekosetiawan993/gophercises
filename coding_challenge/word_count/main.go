package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// filename := "one_word_line.txt"
	args := os.Args

	inputTerminal, _ := os.Stdin.Stat()

	if (inputTerminal.Mode() & os.ModeCharDevice) == 0 {
		inputContentOs, _ := io.ReadAll(os.Stdin)

		numberOfBytes := countBytes(inputContentOs) - 2
		numberOfWords := countWords(inputContentOs)
		numberOfLines := countLines(inputContentOs)
		fmt.Printf("\t%v\t%v\t%v", numberOfLines, numberOfWords, numberOfBytes)

		return
	} else if len(args) < 3 {
		fmt.Printf("go_wc: %v: go_wc need at least three arguments", args)
		os.Exit(1)
	}
	// else {
	// 	inputContent = readFile(args[2])
	// 	filename = args[2]
	// }

	filename := args[len(args)-1]
	inputContent := readFile(filename)

	if args[1] == "-c" {
		numberOfBytes := countBytes(inputContent)
		fmt.Printf("%d %v", numberOfBytes, filename)
	} else if args[1] == "-w" {
		numberOfWords := countWords(inputContent)
		fmt.Printf("%d %v", numberOfWords, filename)
	} else if args[1] == "-l" {
		numberOfLines := countLines(inputContent)
		fmt.Printf("%d %v", numberOfLines, filename)
	} else if args[1] == "-m" {
		numberOfLines := countCharacters(inputContent)
		fmt.Printf("%d %v", numberOfLines, filename)
	} else if args[1] == "-L" {
		maxLineLength := calculateMaxLineLength(inputContent)
		fmt.Printf("%d %v", maxLineLength, filename)
	} else {
		numberOfBytes := countBytes(inputContent)
		numberOfWords := countWords(inputContent)
		numberOfLines := countLines(inputContent)
		fmt.Printf("No such command %v \n", args[1])
		fmt.Printf("%d  %d  %d  %v", numberOfBytes, numberOfLines, numberOfWords, filename)
	}

	// fmt.Printf("The number of bytes is: %v\n", numberOfBytes)
	// fmt.Printf("The number of lines is: %v\n", numberOfLines)
	// fmt.Printf("The number of words is: %v\n", numberOfWords)
	// fmt.Println(string(readFile(filename)))
	// fmt.Println("Here we go again.")
}

func readFile(filename string) []byte {
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("go_wc: %v: No such file or directory", filename)
		os.Exit(1)
	}
	return input
}

func countBytes(content []byte) int64 {
	return int64(len(content))
}

func countLines(content []byte) int64 {
	if len(content) == 0 {
		return int64(0)
	}
	lines := strings.Split(string(content), "\n")

	if lines[len(lines)-1] == "" {
		return int64(len(lines) - 1)
	}
	return int64(len(lines))
}

func countWords(content []byte) int64 {
	if len(content) == 0 {
		return int64(0)
	}

	totalWords := 0

	lines := strings.Split(string(content), "\n")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	for _, line := range lines {
		wordsPerLine := len(strings.Split(line, " "))
		totalWords += wordsPerLine
	}

	return int64(totalWords)
}

func countCharacters(content []byte) int64 {
	return int64(utf8.RuneCount(content))
}

func calculateMaxLineLength(content []byte) int64 {
	if len(content) == 0 {
		return int64(0)
	}
	lines := strings.Split(string(content), "\n")
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}
	return int64(maxLength)
}

// if (inputTerminal.Mode() & os.ModeCharDevice) == 0 {
// 	fmt.Println("input from terminal", inputTerminal.Mode(), os.ModeCharDevice)
// 	inputString, _ := io.ReadAll(os.Stdin)
// 	fmt.Println(string(inputString))
// }
