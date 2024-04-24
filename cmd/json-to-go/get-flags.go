package main

import (
	"flag"
	"fmt"
	"os"
)

func getFlags() (string, string) {
	// get the input flag value
	inputFileFlag := flag.String("input", "", "input json file")
	outputFileFlag := flag.String("output", "output.go", "output go file")

	flag.Parse()

	if *inputFileFlag == "" {
		fmt.Println("input file is required")
		os.Exit(1)
	}

	if *outputFileFlag == "" {
		fmt.Println("output file is required")
		os.Exit(1)
	}

	return *inputFileFlag, *outputFileFlag
}

func checkFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
