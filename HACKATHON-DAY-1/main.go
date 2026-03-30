package main

import (
	"fmt"
)

func main() {

	var num1 float64
	var opp string
	var num2 float64
	var choice string

	for {
		fmt.Print("Enter a number\n")
		_, err := fmt.Scanln(&num1)
		if err != nil {
			fmt.Println("invalid input: Enter a valid number.")

		}

		fmt.Println("chose a opperator: +,/,-,*")
		fmt.Scanln(&opp)

		fmt.Println("Enter the second number")
		_, err1 := fmt.Scanln(&num2)
		if err1 != nil {
			fmt.Println("Invalid input: Enter a valid second number")
			continue
		}

		switch opp {
		case "+":
			fmt.Println("Result:", num1+num2)
		case "-":
			fmt.Println("Result:", num1-num2)
		case "*":
			fmt.Println("Result:", num1*num2)
		case "/":
			fmt.Println("Result:", num1/num2)
			if num2 == 0 {
				fmt.Println("Invalid: num can not be divisible by zero. input a valid number")
			}
		default:
			fmt.Println("Invalid operator, chose between:,  +, /, -, *")

		}
		fmt.Println("wish to continue, type 'yes' else type 'no'")
		fmt.Scanln(&choice)

		if choice == "yes" {
			fmt.Println("you can continue")
			continue
		}
		if choice == "No" {
			fmt.Println("Do have a lovely day ahead")
		}
		if choice != "yes" && choice != "no" {
			fmt.Println("Invalide:, Enter a yes or a no")
			continue
		}

		break
	}

}

