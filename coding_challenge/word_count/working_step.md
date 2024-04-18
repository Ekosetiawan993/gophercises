### Counting words

Here basically I try to count the words by iterating over each line, and separate the line content using space.

```go
func countWords(content []byte) int64 {
	if len(content) == 0 {
		return int64(0)
	}

	totalWords := 0

	lines := strings.Split(string(content), "\n")

    // skip the last line if it is empty
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	for _, line := range lines {
		wordsPerLine := len(strings.Split(line, " "))
		totalWords += wordsPerLine
	}

	return int64(totalWords)
}
```

#### Codes before adding args

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "one_word_line.txt"
	inputContent := readFile(filename)
	numberOfBytes := countBytes(inputContent)
	numberOfLines := countLines(inputContent)
	numberOfWords := countWords(inputContent)

	fmt.Printf("The number of bytes is: %v\n", numberOfBytes)
	fmt.Printf("The number of lines is: %v\n", numberOfLines)
	fmt.Printf("The number of words is: %v\n", numberOfWords)
	// fmt.Println(string(readFile(filename)))
	fmt.Println("Here we go again.")
}

func readFile(filename string) []byte {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
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
		// return int64(len(lines) - 1)
	}

	for _, line := range lines {
		wordsPerLine := len(strings.Split(line, " "))
		totalWords += wordsPerLine
	}

	return int64(totalWords)
}

```
