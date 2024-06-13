package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("hello arithmetic encoding")
	// read read a file
	file, err := os.Args[1:]
	if err != nil {
		log.Fatal(err.New(err))
	}

	// for each file i want to have a freq table for the entries
	for _, file := range os.Args[1:] {
		fmt.Printf("file %f\n", file)
	}

}
