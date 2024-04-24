package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// var nFlag = flag.Int("n", 1234, "help message for flag n")

func main() {
	// get the input flag value
	inputFileFlag := flag.String("input", "", "input json file")
	outputFileFlag := flag.String("output", "output.go", "output go file")

	flag.Parse()

	fmt.Println("input file:", *inputFileFlag)
	fmt.Println("output file:", *outputFileFlag)

	if *inputFileFlag == "" {
		fmt.Println("input file is required")
		os.Exit(1)
	}

	if *outputFileFlag == "" {
		fmt.Println("output file is required")
		os.Exit(1)
	}

	if !checkFileExists(*inputFileFlag) {
		fmt.Println("input file does not exist")
		os.Exit(1)
	}

	if checkFileExists(*outputFileFlag) {
		fmt.Println("output file already exists")
		os.Exit(1)
	}

	// read the input file
	input, err := os.ReadFile(*inputFileFlag)
	if err != nil {
		fmt.Println("error reading input file")
		os.Exit(1)
	}

	// parse the input file
	var data interface{}
	err = json.Unmarshal(input, &data)
	if err != nil {
		fmt.Println("error parsing input JSON:", err)
		os.Exit(1)
	}
	// generate the output file
	fmt.Println(data)
	// write the output file
}

func checkFileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
