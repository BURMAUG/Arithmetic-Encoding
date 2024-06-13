package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello arithmetic encoding")
	// read read a file
	if len(os.Args[1:]) <= 0 {
		log.Fatal(errors.New(" Add files for compression in the command line"))
	}
	// for each file i want to have a freq table for the entries
	for _, file := range os.Args[1:] {
		fmt.Printf("file %f\n", file)
	}

}
