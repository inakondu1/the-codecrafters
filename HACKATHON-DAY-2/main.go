package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Use this format: <convert> <value> <bin|hex|dec>")
		fmt.Println("Type 'stop' to exit")
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		word := strings.Fields(input)

		
		if len(word) == 0 {
			continue
		}

		
		if word[0] == "stop" {
			fmt.Println("Goodbye")
			break
		}

		
		if word[0] != "convert" {
			fmt.Println("Invalid command")
			continue
		}

		
		if len(word) != 3 {
			fmt.Println("Invalid number of arguments")
			continue
		}

		value := word[1]
		mode := word[2]

		switch mode {

		case "hex":
			num, err := strconv.ParseInt(value, 16, 64)
			if err != nil {
				fmt.Println("Invalid hex value")
				continue
			}
			fmt.Println("✦ Decimal:", num)

		case "bin":
			num, err := strconv.ParseInt(value, 2, 64)
			if err != nil {
				fmt.Println("Invalid binary value")
				continue
			}
			fmt.Println("✦ Decimal:", num)

		case "dec":
			num, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				fmt.Println("Invalid decimal value")
				continue
			}
			fmt.Println("✦ Binary:", strconv.FormatInt(num, 2))
			fmt.Println("✦ Hex:", strings.ToUpper(strconv.FormatInt(num, 16)))

		default:
			fmt.Println("Invalid format (use bin, hex, or dec)")
		}
	}
}