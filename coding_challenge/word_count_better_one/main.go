package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var inputContent []byte
	var filename string

	if inputContent, filename = getInput(); len(inputContent) == 0 {
		fmt.Println("how to use: go_wc <flag> <filename> \nMake sure your file is not empty")
		os.Exit(1)
	}
	// already get inputContent in []byte
	flags := GetFlags()

	// initial regular reader
	initializedReaders := InitializeReaders()

	if flagNamePassed, flagHasBeenPassed := HaveBeenPassed(flags); flagHasBeenPassed {
		// if user pass flag
		count := CountWithSpecificReader(initializedReaders[flagNamePassed], inputContent)
		fmt.Printf("\t%d  %s\n", count, filename)
		os.Exit(0)
	}

	bytes, words, lines := CountBytesWordsAndLines(initializedReaders, inputContent)
	fmt.Printf("\t%d\t%d\t%d\t%v", bytes, words, lines, filename)

}

func getInput() ([]byte, string) {
	if HasInputFromTerminal() {
		// if input from terminal
		// fmt.Println("from terminal")
		return ReadInputFromTerminal(), ""
	}

	if HasProvidedArgs() {
		filename := GetFilename()
		// fmt.Println("get filename", filename)
		return readFileWithFilename(filename), filename
	}

	return make([]byte, 0), ""
}

func HasInputFromTerminal() bool {
	// return 0 == 0 / true if input from terminal
	// return 1 == 0 / false if input from file
	f, _ := os.Stdin.Stat()
	return (f.Mode() & os.ModeCharDevice) == 0
}

func ReadInputFromTerminal() []byte {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading file from terminal: %v", err)
	}
	return input
}

func HasProvidedArgs() bool {
	return len(os.Args) > 1
}

func HaveBeenPassed(flags map[string]*bool) (string, bool) {
	for flagName, flagValue := range flags {
		if *flagValue {
			return flagName, true
		}
	}
	return "", false
}

func GetFilename() string {
	return os.Args[len(os.Args)-1]
}

func readFileWithFilename(filename string) []byte {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading text file %v", err)
	}
	return input
}

const (
	BytesFlag = "c"
	LinesFlag = "l"
	WordsFlag = "w"
	CharsFlag = "m"
)

func GetFlags() map[string]*bool {
	flags := map[string]*bool{
		BytesFlag: flag.Bool(BytesFlag, false, "Count bytes"),
		LinesFlag: flag.Bool(LinesFlag, false, "Count lines"),
		WordsFlag: flag.Bool(WordsFlag, false, "Count words"),
		CharsFlag: flag.Bool(CharsFlag, false, "Count chars"),
	}
	flag.Parse()
	return flags
}

type WcReaderInterface interface {
	Count(content []byte) int64
}

func InitializeReaders() map[string]WcReaderInterface {
	return map[string]WcReaderInterface{
		BytesFlag: NewWcBytesReader(),
		LinesFlag: NewWcLinesReader(),
		WordsFlag: NewWcWordsReader(),
		CharsFlag: NewWcCharsReader(),
	}
}

func CountWithSpecificReader(specificReader WcReaderInterface, input []byte) int64 {
	return specificReader.Count(input)
}

func CountBytesWordsAndLines(readers map[string]WcReaderInterface, input []byte) (int64, int64, int64) {
	return readers[BytesFlag].Count(input),
		readers[WordsFlag].Count(input),
		readers[LinesFlag].Count(input)
}

// define readers
type WcBytesReader struct{}

func NewWcBytesReader() WcBytesReader {
	return WcBytesReader{}
}
func (w WcBytesReader) Count(content []byte) int64 {
	return int64(len(content))
}

type WcLinesReader struct{}

func NewWcLinesReader() WcLinesReader {
	return WcLinesReader{}
}
func (w WcLinesReader) Count(content []byte) int64 {
	if len(content) == 0 {
		return int64(0)
	}
	lines := strings.Split(string(content), "\n")

	if lines[len(lines)-1] == "" {
		return int64(len(lines) - 1)
	}
	return int64(len(lines))
}

type WcWordsReader struct{}

func NewWcWordsReader() WcWordsReader {
	return WcWordsReader{}
}
func (w WcWordsReader) Count(content []byte) int64 {
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

type WcCharsReader struct{}

func NewWcCharsReader() WcCharsReader {
	return WcCharsReader{}
}
func (w WcCharsReader) Count(content []byte) int64 {
	return int64(utf8.RuneCount(content))
}
