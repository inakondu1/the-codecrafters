package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . <input><output>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Println("Error: the inputFile and outputFile must be different.")
		return
	}
	text, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error: Failed to read inputFile")
		return
	}
	result := binTodecimal(string(text))
	result = HexToDecimal(result)
	result = capitalizeWord(result)
	result = FixQuotes(result)
	result = punctuation(result)
	result = Article(result)
	result = edit(result)

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error: Failed to write outputFile.")
	}
}
