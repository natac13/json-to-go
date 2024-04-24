package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	// get the input and output file flags
	inputFileFlag, outputFileFlag := getFlags()

	if !checkFileExists(inputFileFlag) {
		fmt.Println("input file does not exist")
		os.Exit(1)
	}

	if checkFileExists(outputFileFlag) {
		// remove the file if it already exists
		err := os.Remove(outputFileFlag)
		if err != nil {
			fmt.Println("error removing existing output file")
			os.Exit(1)
		}
	}

	// read the input file
	input, err := os.ReadFile(inputFileFlag)
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

	jsonToGo, err := NewJsonToGo(input)
	if err != nil {
		os.Exit(1)
	}

	output := jsonToGo.Generate("")

	// write the output file
	fmt.Println("writing output file")
	fmt.Println(output.GoCode)

	output.GoCode = fmt.Sprintf("package main\n\n%s", output.GoCode)

	err = os.WriteFile(outputFileFlag, []byte(output.GoCode), 0644)

	if err != nil {
		fmt.Println("error writing output file")
		os.Exit(1)
	}

	// run goimports on the output file

	// format the output file
	fmt.Println("formatting output file")
	cmd := exec.Command("gofmt", "-w", outputFileFlag)
	err = cmd.Run()

	if err != nil {
		fmt.Println("error formatting output file")
		os.Exit(1)
	}

	fmt.Println("running goimports on output file")
	cmd = exec.Command("gopls", "imports", "-w", outputFileFlag)
	err = cmd.Run()

	if err != nil {
		fmt.Println("error running goimports on output file")
		os.Exit(1)
	}
}
