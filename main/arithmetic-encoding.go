package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	var fileContent string
	runeFreqMap := make(map[rune]int, 10) // rune freq map
	fmt.Println("hello arithmetic encoding")

	// Check for a file being read from the console
	if len(os.Args[1:]) <= 0 {
		log.Fatal(errors.New("Add files for compression in the command line"))
	}

	// for each file i want to have a freq table for the entries
	for _, file := range os.Args[1:] {
		fmt.Printf("file %f\n", file)
		filem, err := os.ReadFile(file)
		if err != nil {
			panic("No file exist")
		}
		fileContent = string(filem)
	}

	for _, d := range fileContent {
		//	fmt.Printf("%c", d)
		runeFreqMap[d]++
	}

	for k, v := range runeFreqMap {
		relativeFrequency := float64(v) / float64(len(fileContent))
		fmt.Printf("%q: %d, Relative Frequency: %.6f\n", k, v, relativeFrequency)
	}
}
