package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(word string) string {
	if len(word) == 0 {
		return ""
	}

	return reverse(word[1:]) + word[:1]
}

func snakecase(word string) string {
	return strings.ReplaceAll(word, " ", "_")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter a word")

		input, _ := reader.ReadString('\n')
		word := strings.Fields(input)

		var Transformation string

		fmt.Scanln(&word)

		if len(word) == 0 {
			fmt.Println("no text provided")
			continue
		}

		fmt.Println("chose a Transformation:")
		fmt.Println("<1> Toupper")
		fmt.Println("<2> Tolower")
		fmt.Println("<3> Fields")
		fmt.Println("<4> Title")
		fmt.Println("<5> Snake")
		fmt.Println("<6> Reverse")
		fmt.Println("<7> exit")

		fmt.Scanln(&Transformation)

		switch Transformation {
		case "1":
			result := word
			input := strings.Join(result, " ")
			fmt.Println("Result:", strings.ToUpper(input))
		case "2":
			result := word
			input := strings.Join(result, " ")

			fmt.Println("Return:", strings.ToLower(input))
		case "3":
			result := word
			input := strings.Join(result, " ")
			fmt.Println("Return:", strings.Fields(input))
		case "4":
			result := word
			input := strings.Join(result, " ")
			fmt.Println("Return:", strings.Title(input))
		case "5":
			result := word
			input := strings.Join(result, " ")
			result1 := strings.ToLower(input)
			fmt.Println(snakecase(result1))

		case "6":
			good := strings.Fields(input)
			for i := 0; i < len(good); i++ {
				good[i] = reverse(good[i])
			}
			fmt.Println(strings.Join(good, " "))
		case "7":
			fmt.Println("goodbye")
			return
		default:
			fmt.Println("invalide input")

		}
	}
}
