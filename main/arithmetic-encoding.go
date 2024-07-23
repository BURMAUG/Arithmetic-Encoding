package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	fileContent string
	codeTable   []CodeProbabilityTable
)

// Encoder represents
type Encoder interface {
	EventEncoder([]byte) string // maybe takes a file and return the encoding
}

// CodeProbabilityTable represents the idea of what the probability table
// for the string data
type CodeProbabilityTable struct {
	symbol      string
	freq        int
	probability float64
}

// The more frequent symbols have shorter codes
// The less frequent symbols have longer codes

// I need a way to map symbols to their respective codes
func main() {
	runeFreqMap := make(map[rune]int, 0) // rune freq map
	fmt.Println("hello arithmetic encoding")

	// Check for a file being read from the console
	if len(os.Args[1:]) <= 0 {
		log.Fatal(errors.New("Add files for compression in the command line"))
	}

	// for each file i want to have a freq table for the entries
	for _, file := range os.Args[1:] {
		filem, err := os.ReadFile(file)
		if err != nil {
			panic("No file exist")
		}
		fileContent = string(filem)
	}

	// makes the frequency table
	MakeFreqTable(runeFreqMap)
	// add the obj
	for k, val := range runeFreqMap {
		relativeProb := float64(val) / float64(len(fileContent))
		codeTable = append(codeTable, CodeProbabilityTable{
			string(k),
			val,
			relativeProb,
		})
	}
	// prints the code struct
	for _, codes := range codeTable {
		fmt.Printf("%q %d %.3f\n", codes.symbol, codes.freq, codes.probability)
	}
}

// Todo make Preqencies table here
func MakeFreqTable(data map[rune]int) {
	for _, d := range fileContent {
		data[d]++
	}
}

// Todo

// Todo
